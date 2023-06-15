package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisPoint(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_point")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_point

	for rows.Next() {
		var seis_point dto.Seis_point
		if err := rows.Scan(&seis_point.Seis_set_subtype, &seis_point.Seis_set_id, &seis_point.Seis_point_id, &seis_point.Acqtn_seq_no, &seis_point.Active_ind, &seis_point.Bend_ind, &seis_point.Datum_elev, &seis_point.Datum_elev_ouom, &seis_point.Effective_date, &seis_point.Elev, &seis_point.Elev_ouom, &seis_point.End_ind, &seis_point.Exception_ind, &seis_point.Expiry_date, &seis_point.Flowing_hole_ind, &seis_point.Mapping_code, &seis_point.Measured_depth, &seis_point.Measured_depth_ouom, &seis_point.Ppdm_guid, &seis_point.Reference_datum, &seis_point.Remark, &seis_point.Seis_point_label, &seis_point.Seis_point_lat, &seis_point.Seis_point_long, &seis_point.Seis_point_no, &seis_point.Seis_station_type, &seis_point.Source, &seis_point.Spatial_seq_no, &seis_point.X_coordinate, &seis_point.Y_coordinate, &seis_point.Row_changed_by, &seis_point.Row_changed_date, &seis_point.Row_created_by, &seis_point.Row_created_date, &seis_point.Row_effective_date, &seis_point.Row_expiry_date, &seis_point.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_point to result
		result = append(result, seis_point)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisPoint(c *fiber.Ctx) error {
	var seis_point dto.Seis_point

	setDefaults(&seis_point)

	if err := json.Unmarshal(c.Body(), &seis_point); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_point values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	seis_point.Row_created_date = formatDate(seis_point.Row_created_date)
	seis_point.Row_changed_date = formatDate(seis_point.Row_changed_date)
	seis_point.Effective_date = formatDateString(seis_point.Effective_date)
	seis_point.Expiry_date = formatDateString(seis_point.Expiry_date)
	seis_point.Row_effective_date = formatDateString(seis_point.Row_effective_date)
	seis_point.Row_expiry_date = formatDateString(seis_point.Row_expiry_date)






	rows, err := stmt.Exec(seis_point.Seis_set_subtype, seis_point.Seis_set_id, seis_point.Seis_point_id, seis_point.Acqtn_seq_no, seis_point.Active_ind, seis_point.Bend_ind, seis_point.Datum_elev, seis_point.Datum_elev_ouom, seis_point.Effective_date, seis_point.Elev, seis_point.Elev_ouom, seis_point.End_ind, seis_point.Exception_ind, seis_point.Expiry_date, seis_point.Flowing_hole_ind, seis_point.Mapping_code, seis_point.Measured_depth, seis_point.Measured_depth_ouom, seis_point.Ppdm_guid, seis_point.Reference_datum, seis_point.Remark, seis_point.Seis_point_label, seis_point.Seis_point_lat, seis_point.Seis_point_long, seis_point.Seis_point_no, seis_point.Seis_station_type, seis_point.Source, seis_point.Spatial_seq_no, seis_point.X_coordinate, seis_point.Y_coordinate, seis_point.Row_changed_by, seis_point.Row_changed_date, seis_point.Row_created_by, seis_point.Row_created_date, seis_point.Row_effective_date, seis_point.Row_expiry_date, seis_point.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisPoint(c *fiber.Ctx) error {
	var seis_point dto.Seis_point

	setDefaults(&seis_point)

	if err := json.Unmarshal(c.Body(), &seis_point); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_point.Seis_set_subtype = id

    if seis_point.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_point set seis_set_id = :1, seis_point_id = :2, acqtn_seq_no = :3, active_ind = :4, bend_ind = :5, datum_elev = :6, datum_elev_ouom = :7, effective_date = :8, elev = :9, elev_ouom = :10, end_ind = :11, exception_ind = :12, expiry_date = :13, flowing_hole_ind = :14, mapping_code = :15, measured_depth = :16, measured_depth_ouom = :17, ppdm_guid = :18, reference_datum = :19, remark = :20, seis_point_label = :21, seis_point_lat = :22, seis_point_long = :23, seis_point_no = :24, seis_station_type = :25, source = :26, spatial_seq_no = :27, x_coordinate = :28, y_coordinate = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where seis_set_subtype = :37")
	if err != nil {
		return err
	}

	seis_point.Row_changed_date = formatDate(seis_point.Row_changed_date)
	seis_point.Effective_date = formatDateString(seis_point.Effective_date)
	seis_point.Expiry_date = formatDateString(seis_point.Expiry_date)
	seis_point.Row_effective_date = formatDateString(seis_point.Row_effective_date)
	seis_point.Row_expiry_date = formatDateString(seis_point.Row_expiry_date)






	rows, err := stmt.Exec(seis_point.Seis_set_id, seis_point.Seis_point_id, seis_point.Acqtn_seq_no, seis_point.Active_ind, seis_point.Bend_ind, seis_point.Datum_elev, seis_point.Datum_elev_ouom, seis_point.Effective_date, seis_point.Elev, seis_point.Elev_ouom, seis_point.End_ind, seis_point.Exception_ind, seis_point.Expiry_date, seis_point.Flowing_hole_ind, seis_point.Mapping_code, seis_point.Measured_depth, seis_point.Measured_depth_ouom, seis_point.Ppdm_guid, seis_point.Reference_datum, seis_point.Remark, seis_point.Seis_point_label, seis_point.Seis_point_lat, seis_point.Seis_point_long, seis_point.Seis_point_no, seis_point.Seis_station_type, seis_point.Source, seis_point.Spatial_seq_no, seis_point.X_coordinate, seis_point.Y_coordinate, seis_point.Row_changed_by, seis_point.Row_changed_date, seis_point.Row_created_by, seis_point.Row_effective_date, seis_point.Row_expiry_date, seis_point.Row_quality, seis_point.Seis_set_subtype)
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

func PatchSeisPoint(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_point set "
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
	query += " where seis_set_subtype = :id"

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

func DeleteSeisPoint(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_point dto.Seis_point
	seis_point.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_point where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_point.Seis_set_subtype)
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


