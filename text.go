var dateFormat = "%Y-%m-%d"

selectQuery := fmt.Sprintf(`
	WITH raw_occupied_available_room AS (
		SELECT 
			spn.date as stay_date, count(spn.reservationId) as occupied_room, (SELECT 
					COUNT(r.id)
				FROM
					Rooms r
						INNER JOIN
					RoomTypeDetails rtd ON rtd.roomType = r.roomTypeId
						AND rtd.hotelId = r.hotelId
				WHERE
					r.hotelId = %d AND r.isRemoved = 0
						AND rtd.isActive = TRUE
						AND rtd.roomType NOT IN (5 , 6)) total_room
		FROM
			BbxStayPerNights spn
				INNER JOIN
			BbxReservations r ON r.id = spn.reservationId
				INNER JOIN
			BbxStays s ON s.id = spn.stayId
		WHERE
		   spn.date BETWEEN '%s' AND '%s'
		   AND r.hotelId = %d
		   AND r.paymentStatusId = 2
		   AND r.isRemoved = FALSE
		   AND r.markAsDeleted = FALSE
		GROUP BY 1 ORDER BY 1 ASC),
		raw_ooo_room AS (
		SELECT 
			DATE_FORMAT(roo.date, '%s') ooo_date,
			COUNT(r.id) ooo_room
		FROM
			Rooms r
				INNER JOIN
			RoomTypeDetails rtd ON rtd.roomType = r.roomTypeId
				AND rtd.hotelId = r.hotelId
				INNER JOIN
			BbxRoomStatusSpareAndOutOrders roo ON roo.roomId = r.id
		WHERE
			r.hotelId = %d AND r.isRemoved = 0
				AND rtd.isActive = TRUE
				AND rtd.roomType NOT IN (5 , 6)
				AND roo.status != 'SPARE'
				AND roo.date >= '%s'
				AND roo.date <= '%s'
		GROUP BY 1
		),
		total_available_room AS (
		SELECT 
			roar.stay_date, total_room - IF(ooo_room IS NULL, 0, ooo_room) - occupied_room AS total_available_room
		FROM 
			raw_occupied_available_room roar 
		LEFT JOIN 
			raw_ooo_room ror 
				ON roar.stay_date = ror.ooo_date)
		SELECT 
			DATE_FORMAT(stay_date, '%s') as soldout_date 
		FROM 
			total_available_room 
		WHERE 
			total_available_room <= 0
	`, hotelID, firstDate, lastDate, hotelID, dateFormat, hotelID, firstDate, lastDate, dateFormat)