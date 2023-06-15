package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpLinePoint(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_line_point")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_line_point

	for rows.Next() {
		var sp_line_point dto.Sp_line_point
		if err := rows.Scan(&sp_line_point.Line_id, &sp_line_point.Point_seq_no, &sp_line_point.Active_ind, &sp_line_point.Bend_ind, &sp_line_point.Depth, &sp_line_point.Depth_ouom, &sp_line_point.Effective_date, &sp_line_point.Elevation, &sp_line_point.Elevation_ouom, &sp_line_point.Expiry_date, &sp_line_point.First_point_ind, &sp_line_point.Last_point_ind, &sp_line_point.Latitude, &sp_line_point.Location_quality, &sp_line_point.Longitude, &sp_line_point.Point_label, &sp_line_point.Point_no, &sp_line_point.Point_position, &sp_line_point.Ppdm_guid, &sp_line_point.Remark, &sp_line_point.Source, &sp_line_point.Row_changed_by, &sp_line_point.Row_changed_date, &sp_line_point.Row_created_by, &sp_line_point.Row_created_date, &sp_line_point.Row_effective_date, &sp_line_point.Row_expiry_date, &sp_line_point.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_line_point to result
		result = append(result, sp_line_point)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpLinePoint(c *fiber.Ctx) error {
	var sp_line_point dto.Sp_line_point

	setDefaults(&sp_line_point)

	if err := json.Unmarshal(c.Body(), &sp_line_point); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_line_point values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	sp_line_point.Row_created_date = formatDate(sp_line_point.Row_created_date)
	sp_line_point.Row_changed_date = formatDate(sp_line_point.Row_changed_date)
	sp_line_point.Effective_date = formatDateString(sp_line_point.Effective_date)
	sp_line_point.Expiry_date = formatDateString(sp_line_point.Expiry_date)
	sp_line_point.Row_effective_date = formatDateString(sp_line_point.Row_effective_date)
	sp_line_point.Row_expiry_date = formatDateString(sp_line_point.Row_expiry_date)






	rows, err := stmt.Exec(sp_line_point.Line_id, sp_line_point.Point_seq_no, sp_line_point.Active_ind, sp_line_point.Bend_ind, sp_line_point.Depth, sp_line_point.Depth_ouom, sp_line_point.Effective_date, sp_line_point.Elevation, sp_line_point.Elevation_ouom, sp_line_point.Expiry_date, sp_line_point.First_point_ind, sp_line_point.Last_point_ind, sp_line_point.Latitude, sp_line_point.Location_quality, sp_line_point.Longitude, sp_line_point.Point_label, sp_line_point.Point_no, sp_line_point.Point_position, sp_line_point.Ppdm_guid, sp_line_point.Remark, sp_line_point.Source, sp_line_point.Row_changed_by, sp_line_point.Row_changed_date, sp_line_point.Row_created_by, sp_line_point.Row_created_date, sp_line_point.Row_effective_date, sp_line_point.Row_expiry_date, sp_line_point.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpLinePoint(c *fiber.Ctx) error {
	var sp_line_point dto.Sp_line_point

	setDefaults(&sp_line_point)

	if err := json.Unmarshal(c.Body(), &sp_line_point); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_line_point.Line_id = id

    if sp_line_point.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_line_point set point_seq_no = :1, active_ind = :2, bend_ind = :3, depth = :4, depth_ouom = :5, effective_date = :6, elevation = :7, elevation_ouom = :8, expiry_date = :9, first_point_ind = :10, last_point_ind = :11, latitude = :12, location_quality = :13, longitude = :14, point_label = :15, point_no = :16, point_position = :17, ppdm_guid = :18, remark = :19, source = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where line_id = :28")
	if err != nil {
		return err
	}

	sp_line_point.Row_changed_date = formatDate(sp_line_point.Row_changed_date)
	sp_line_point.Effective_date = formatDateString(sp_line_point.Effective_date)
	sp_line_point.Expiry_date = formatDateString(sp_line_point.Expiry_date)
	sp_line_point.Row_effective_date = formatDateString(sp_line_point.Row_effective_date)
	sp_line_point.Row_expiry_date = formatDateString(sp_line_point.Row_expiry_date)






	rows, err := stmt.Exec(sp_line_point.Point_seq_no, sp_line_point.Active_ind, sp_line_point.Bend_ind, sp_line_point.Depth, sp_line_point.Depth_ouom, sp_line_point.Effective_date, sp_line_point.Elevation, sp_line_point.Elevation_ouom, sp_line_point.Expiry_date, sp_line_point.First_point_ind, sp_line_point.Last_point_ind, sp_line_point.Latitude, sp_line_point.Location_quality, sp_line_point.Longitude, sp_line_point.Point_label, sp_line_point.Point_no, sp_line_point.Point_position, sp_line_point.Ppdm_guid, sp_line_point.Remark, sp_line_point.Source, sp_line_point.Row_changed_by, sp_line_point.Row_changed_date, sp_line_point.Row_created_by, sp_line_point.Row_effective_date, sp_line_point.Row_expiry_date, sp_line_point.Row_quality, sp_line_point.Line_id)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchSpLinePoint(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_line_point set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where line_id = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteSpLinePoint(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_line_point dto.Sp_line_point
	sp_line_point.Line_id = id

	stmt, err := db.Prepare("delete from sp_line_point where line_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_line_point.Line_id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


