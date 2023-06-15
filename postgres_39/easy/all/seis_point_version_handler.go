package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisPointVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_point_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_point_version

	for rows.Next() {
		var seis_point_version dto.Seis_point_version
		if err := rows.Scan(&seis_point_version.Seis_set_subtype, &seis_point_version.Seis_line_id, &seis_point_version.Seis_point_id, &seis_point_version.Seis_point_obs_no, &seis_point_version.Acquisition_id, &seis_point_version.Active_ind, &seis_point_version.Easting, &seis_point_version.Easting_ouom, &seis_point_version.Effective_date, &seis_point_version.Elev, &seis_point_version.Elev_ouom, &seis_point_version.Ew_direction, &seis_point_version.Ew_distance, &seis_point_version.Ew_start_line, &seis_point_version.Expiry_date, &seis_point_version.Geog_coord_system_id, &seis_point_version.Local_coord_system_id, &seis_point_version.Map_coord_system_id, &seis_point_version.Northing, &seis_point_version.Northing_ouom, &seis_point_version.North_type, &seis_point_version.Ns_direction, &seis_point_version.Ns_distance, &seis_point_version.Ns_start_line, &seis_point_version.Polar_azimuth, &seis_point_version.Polar_offset, &seis_point_version.Polar_offset_ouom, &seis_point_version.Ppdm_guid, &seis_point_version.Preferred_ind, &seis_point_version.Reference_loc, &seis_point_version.Remark, &seis_point_version.Seis_point_label, &seis_point_version.Seis_point_lat, &seis_point_version.Seis_point_long, &seis_point_version.Seis_point_no, &seis_point_version.Source, &seis_point_version.Uwi, &seis_point_version.Version_type, &seis_point_version.X_coordinate, &seis_point_version.Y_coordinate, &seis_point_version.Row_changed_by, &seis_point_version.Row_changed_date, &seis_point_version.Row_created_by, &seis_point_version.Row_created_date, &seis_point_version.Row_effective_date, &seis_point_version.Row_expiry_date, &seis_point_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_point_version to result
		result = append(result, seis_point_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisPointVersion(c *fiber.Ctx) error {
	var seis_point_version dto.Seis_point_version

	setDefaults(&seis_point_version)

	if err := json.Unmarshal(c.Body(), &seis_point_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_point_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47)")
	if err != nil {
		return err
	}
	seis_point_version.Row_created_date = formatDate(seis_point_version.Row_created_date)
	seis_point_version.Row_changed_date = formatDate(seis_point_version.Row_changed_date)
	seis_point_version.Effective_date = formatDateString(seis_point_version.Effective_date)
	seis_point_version.Expiry_date = formatDateString(seis_point_version.Expiry_date)
	seis_point_version.Row_effective_date = formatDateString(seis_point_version.Row_effective_date)
	seis_point_version.Row_expiry_date = formatDateString(seis_point_version.Row_expiry_date)






	rows, err := stmt.Exec(seis_point_version.Seis_set_subtype, seis_point_version.Seis_line_id, seis_point_version.Seis_point_id, seis_point_version.Seis_point_obs_no, seis_point_version.Acquisition_id, seis_point_version.Active_ind, seis_point_version.Easting, seis_point_version.Easting_ouom, seis_point_version.Effective_date, seis_point_version.Elev, seis_point_version.Elev_ouom, seis_point_version.Ew_direction, seis_point_version.Ew_distance, seis_point_version.Ew_start_line, seis_point_version.Expiry_date, seis_point_version.Geog_coord_system_id, seis_point_version.Local_coord_system_id, seis_point_version.Map_coord_system_id, seis_point_version.Northing, seis_point_version.Northing_ouom, seis_point_version.North_type, seis_point_version.Ns_direction, seis_point_version.Ns_distance, seis_point_version.Ns_start_line, seis_point_version.Polar_azimuth, seis_point_version.Polar_offset, seis_point_version.Polar_offset_ouom, seis_point_version.Ppdm_guid, seis_point_version.Preferred_ind, seis_point_version.Reference_loc, seis_point_version.Remark, seis_point_version.Seis_point_label, seis_point_version.Seis_point_lat, seis_point_version.Seis_point_long, seis_point_version.Seis_point_no, seis_point_version.Source, seis_point_version.Uwi, seis_point_version.Version_type, seis_point_version.X_coordinate, seis_point_version.Y_coordinate, seis_point_version.Row_changed_by, seis_point_version.Row_changed_date, seis_point_version.Row_created_by, seis_point_version.Row_created_date, seis_point_version.Row_effective_date, seis_point_version.Row_expiry_date, seis_point_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisPointVersion(c *fiber.Ctx) error {
	var seis_point_version dto.Seis_point_version

	setDefaults(&seis_point_version)

	if err := json.Unmarshal(c.Body(), &seis_point_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_point_version.Seis_set_subtype = id

    if seis_point_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_point_version set seis_line_id = :1, seis_point_id = :2, seis_point_obs_no = :3, acquisition_id = :4, active_ind = :5, easting = :6, easting_ouom = :7, effective_date = :8, elev = :9, elev_ouom = :10, ew_direction = :11, ew_distance = :12, ew_start_line = :13, expiry_date = :14, geog_coord_system_id = :15, local_coord_system_id = :16, map_coord_system_id = :17, northing = :18, northing_ouom = :19, north_type = :20, ns_direction = :21, ns_distance = :22, ns_start_line = :23, polar_azimuth = :24, polar_offset = :25, polar_offset_ouom = :26, ppdm_guid = :27, preferred_ind = :28, reference_loc = :29, remark = :30, seis_point_label = :31, seis_point_lat = :32, seis_point_long = :33, seis_point_no = :34, source = :35, uwi = :36, version_type = :37, x_coordinate = :38, y_coordinate = :39, row_changed_by = :40, row_changed_date = :41, row_created_by = :42, row_effective_date = :43, row_expiry_date = :44, row_quality = :45 where seis_set_subtype = :47")
	if err != nil {
		return err
	}

	seis_point_version.Row_changed_date = formatDate(seis_point_version.Row_changed_date)
	seis_point_version.Effective_date = formatDateString(seis_point_version.Effective_date)
	seis_point_version.Expiry_date = formatDateString(seis_point_version.Expiry_date)
	seis_point_version.Row_effective_date = formatDateString(seis_point_version.Row_effective_date)
	seis_point_version.Row_expiry_date = formatDateString(seis_point_version.Row_expiry_date)






	rows, err := stmt.Exec(seis_point_version.Seis_line_id, seis_point_version.Seis_point_id, seis_point_version.Seis_point_obs_no, seis_point_version.Acquisition_id, seis_point_version.Active_ind, seis_point_version.Easting, seis_point_version.Easting_ouom, seis_point_version.Effective_date, seis_point_version.Elev, seis_point_version.Elev_ouom, seis_point_version.Ew_direction, seis_point_version.Ew_distance, seis_point_version.Ew_start_line, seis_point_version.Expiry_date, seis_point_version.Geog_coord_system_id, seis_point_version.Local_coord_system_id, seis_point_version.Map_coord_system_id, seis_point_version.Northing, seis_point_version.Northing_ouom, seis_point_version.North_type, seis_point_version.Ns_direction, seis_point_version.Ns_distance, seis_point_version.Ns_start_line, seis_point_version.Polar_azimuth, seis_point_version.Polar_offset, seis_point_version.Polar_offset_ouom, seis_point_version.Ppdm_guid, seis_point_version.Preferred_ind, seis_point_version.Reference_loc, seis_point_version.Remark, seis_point_version.Seis_point_label, seis_point_version.Seis_point_lat, seis_point_version.Seis_point_long, seis_point_version.Seis_point_no, seis_point_version.Source, seis_point_version.Uwi, seis_point_version.Version_type, seis_point_version.X_coordinate, seis_point_version.Y_coordinate, seis_point_version.Row_changed_by, seis_point_version.Row_changed_date, seis_point_version.Row_created_by, seis_point_version.Row_effective_date, seis_point_version.Row_expiry_date, seis_point_version.Row_quality, seis_point_version.Seis_set_subtype)
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

func PatchSeisPointVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_point_version set "
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

func DeleteSeisPointVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_point_version dto.Seis_point_version
	seis_point_version.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_point_version where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_point_version.Seis_set_subtype)
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


