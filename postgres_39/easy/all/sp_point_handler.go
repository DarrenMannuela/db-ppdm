package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpPoint(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_point")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_point

	for rows.Next() {
		var sp_point dto.Sp_point
		if err := rows.Scan(&sp_point.Point_id, &sp_point.Active_ind, &sp_point.Coord_acquisition_id, &sp_point.Coord_system_id, &sp_point.Datum_elev, &sp_point.Datum_elev_ouom, &sp_point.Effective_date, &sp_point.Elevation, &sp_point.Elevation_ouom, &sp_point.Expiry_date, &sp_point.Latitude, &sp_point.Local_coord_system_id, &sp_point.Location_quality, &sp_point.Location_type, &sp_point.Longitude, &sp_point.Measured_depth, &sp_point.Measured_depth_ouom, &sp_point.Point_label, &sp_point.Point_no, &sp_point.Point_position, &sp_point.Point_seq_no, &sp_point.Ppdm_guid, &sp_point.Reference_datum, &sp_point.Remark, &sp_point.Source, &sp_point.Spatial_description_id, &sp_point.Spatial_obs_no, &sp_point.Row_changed_by, &sp_point.Row_changed_date, &sp_point.Row_created_by, &sp_point.Row_created_date, &sp_point.Row_effective_date, &sp_point.Row_expiry_date, &sp_point.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_point to result
		result = append(result, sp_point)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpPoint(c *fiber.Ctx) error {
	var sp_point dto.Sp_point

	setDefaults(&sp_point)

	if err := json.Unmarshal(c.Body(), &sp_point); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_point values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	sp_point.Row_created_date = formatDate(sp_point.Row_created_date)
	sp_point.Row_changed_date = formatDate(sp_point.Row_changed_date)
	sp_point.Effective_date = formatDateString(sp_point.Effective_date)
	sp_point.Expiry_date = formatDateString(sp_point.Expiry_date)
	sp_point.Row_effective_date = formatDateString(sp_point.Row_effective_date)
	sp_point.Row_expiry_date = formatDateString(sp_point.Row_expiry_date)






	rows, err := stmt.Exec(sp_point.Point_id, sp_point.Active_ind, sp_point.Coord_acquisition_id, sp_point.Coord_system_id, sp_point.Datum_elev, sp_point.Datum_elev_ouom, sp_point.Effective_date, sp_point.Elevation, sp_point.Elevation_ouom, sp_point.Expiry_date, sp_point.Latitude, sp_point.Local_coord_system_id, sp_point.Location_quality, sp_point.Location_type, sp_point.Longitude, sp_point.Measured_depth, sp_point.Measured_depth_ouom, sp_point.Point_label, sp_point.Point_no, sp_point.Point_position, sp_point.Point_seq_no, sp_point.Ppdm_guid, sp_point.Reference_datum, sp_point.Remark, sp_point.Source, sp_point.Spatial_description_id, sp_point.Spatial_obs_no, sp_point.Row_changed_by, sp_point.Row_changed_date, sp_point.Row_created_by, sp_point.Row_created_date, sp_point.Row_effective_date, sp_point.Row_expiry_date, sp_point.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpPoint(c *fiber.Ctx) error {
	var sp_point dto.Sp_point

	setDefaults(&sp_point)

	if err := json.Unmarshal(c.Body(), &sp_point); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_point.Point_id = id

    if sp_point.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_point set active_ind = :1, coord_acquisition_id = :2, coord_system_id = :3, datum_elev = :4, datum_elev_ouom = :5, effective_date = :6, elevation = :7, elevation_ouom = :8, expiry_date = :9, latitude = :10, local_coord_system_id = :11, location_quality = :12, location_type = :13, longitude = :14, measured_depth = :15, measured_depth_ouom = :16, point_label = :17, point_no = :18, point_position = :19, point_seq_no = :20, ppdm_guid = :21, reference_datum = :22, remark = :23, source = :24, spatial_description_id = :25, spatial_obs_no = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where point_id = :34")
	if err != nil {
		return err
	}

	sp_point.Row_changed_date = formatDate(sp_point.Row_changed_date)
	sp_point.Effective_date = formatDateString(sp_point.Effective_date)
	sp_point.Expiry_date = formatDateString(sp_point.Expiry_date)
	sp_point.Row_effective_date = formatDateString(sp_point.Row_effective_date)
	sp_point.Row_expiry_date = formatDateString(sp_point.Row_expiry_date)






	rows, err := stmt.Exec(sp_point.Active_ind, sp_point.Coord_acquisition_id, sp_point.Coord_system_id, sp_point.Datum_elev, sp_point.Datum_elev_ouom, sp_point.Effective_date, sp_point.Elevation, sp_point.Elevation_ouom, sp_point.Expiry_date, sp_point.Latitude, sp_point.Local_coord_system_id, sp_point.Location_quality, sp_point.Location_type, sp_point.Longitude, sp_point.Measured_depth, sp_point.Measured_depth_ouom, sp_point.Point_label, sp_point.Point_no, sp_point.Point_position, sp_point.Point_seq_no, sp_point.Ppdm_guid, sp_point.Reference_datum, sp_point.Remark, sp_point.Source, sp_point.Spatial_description_id, sp_point.Spatial_obs_no, sp_point.Row_changed_by, sp_point.Row_changed_date, sp_point.Row_created_by, sp_point.Row_effective_date, sp_point.Row_expiry_date, sp_point.Row_quality, sp_point.Point_id)
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

func PatchSpPoint(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_point set "
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
	query += " where point_id = :id"

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

func DeleteSpPoint(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_point dto.Sp_point
	sp_point.Point_id = id

	stmt, err := db.Prepare("delete from sp_point where point_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_point.Point_id)
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


