package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisBinPointVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_bin_point_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_bin_point_version

	for rows.Next() {
		var seis_bin_point_version dto.Seis_bin_point_version
		if err := rows.Scan(&seis_bin_point_version.Seis_set_subtype, &seis_bin_point_version.Seis_set_id, &seis_bin_point_version.Bin_grid_id, &seis_bin_point_version.Bin_grid_source, &seis_bin_point_version.Bin_point_id, &seis_bin_point_version.Bin_point_obs_no, &seis_bin_point_version.Acquisition_id, &seis_bin_point_version.Active_ind, &seis_bin_point_version.Easting, &seis_bin_point_version.Easting_ouom, &seis_bin_point_version.Effective_date, &seis_bin_point_version.Elev, &seis_bin_point_version.Elev_ouom, &seis_bin_point_version.Ew_direction, &seis_bin_point_version.Ew_distance, &seis_bin_point_version.Ew_start_line, &seis_bin_point_version.Expiry_date, &seis_bin_point_version.Geog_coord_system_id, &seis_bin_point_version.Local_coord_system_id, &seis_bin_point_version.Map_coord_system_id, &seis_bin_point_version.Northing, &seis_bin_point_version.Northing_ouom, &seis_bin_point_version.North_type, &seis_bin_point_version.Ns_direction, &seis_bin_point_version.Ns_distance, &seis_bin_point_version.Ns_start_line, &seis_bin_point_version.Polar_azimuth, &seis_bin_point_version.Polar_offset, &seis_bin_point_version.Polar_offset_ouom, &seis_bin_point_version.Ppdm_guid, &seis_bin_point_version.Preferred_ind, &seis_bin_point_version.Reference_loc, &seis_bin_point_version.Remark, &seis_bin_point_version.Seis_point_label, &seis_bin_point_version.Seis_point_lat, &seis_bin_point_version.Seis_point_long, &seis_bin_point_version.Seis_point_no, &seis_bin_point_version.Uwi, &seis_bin_point_version.Version_type, &seis_bin_point_version.X_coordinate, &seis_bin_point_version.Y_coordinate, &seis_bin_point_version.Row_changed_by, &seis_bin_point_version.Row_changed_date, &seis_bin_point_version.Row_created_by, &seis_bin_point_version.Row_created_date, &seis_bin_point_version.Row_effective_date, &seis_bin_point_version.Row_expiry_date, &seis_bin_point_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_bin_point_version to result
		result = append(result, seis_bin_point_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisBinPointVersion(c *fiber.Ctx) error {
	var seis_bin_point_version dto.Seis_bin_point_version

	setDefaults(&seis_bin_point_version)

	if err := json.Unmarshal(c.Body(), &seis_bin_point_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_bin_point_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48)")
	if err != nil {
		return err
	}
	seis_bin_point_version.Row_created_date = formatDate(seis_bin_point_version.Row_created_date)
	seis_bin_point_version.Row_changed_date = formatDate(seis_bin_point_version.Row_changed_date)
	seis_bin_point_version.Effective_date = formatDateString(seis_bin_point_version.Effective_date)
	seis_bin_point_version.Expiry_date = formatDateString(seis_bin_point_version.Expiry_date)
	seis_bin_point_version.Row_effective_date = formatDateString(seis_bin_point_version.Row_effective_date)
	seis_bin_point_version.Row_expiry_date = formatDateString(seis_bin_point_version.Row_expiry_date)






	rows, err := stmt.Exec(seis_bin_point_version.Seis_set_subtype, seis_bin_point_version.Seis_set_id, seis_bin_point_version.Bin_grid_id, seis_bin_point_version.Bin_grid_source, seis_bin_point_version.Bin_point_id, seis_bin_point_version.Bin_point_obs_no, seis_bin_point_version.Acquisition_id, seis_bin_point_version.Active_ind, seis_bin_point_version.Easting, seis_bin_point_version.Easting_ouom, seis_bin_point_version.Effective_date, seis_bin_point_version.Elev, seis_bin_point_version.Elev_ouom, seis_bin_point_version.Ew_direction, seis_bin_point_version.Ew_distance, seis_bin_point_version.Ew_start_line, seis_bin_point_version.Expiry_date, seis_bin_point_version.Geog_coord_system_id, seis_bin_point_version.Local_coord_system_id, seis_bin_point_version.Map_coord_system_id, seis_bin_point_version.Northing, seis_bin_point_version.Northing_ouom, seis_bin_point_version.North_type, seis_bin_point_version.Ns_direction, seis_bin_point_version.Ns_distance, seis_bin_point_version.Ns_start_line, seis_bin_point_version.Polar_azimuth, seis_bin_point_version.Polar_offset, seis_bin_point_version.Polar_offset_ouom, seis_bin_point_version.Ppdm_guid, seis_bin_point_version.Preferred_ind, seis_bin_point_version.Reference_loc, seis_bin_point_version.Remark, seis_bin_point_version.Seis_point_label, seis_bin_point_version.Seis_point_lat, seis_bin_point_version.Seis_point_long, seis_bin_point_version.Seis_point_no, seis_bin_point_version.Uwi, seis_bin_point_version.Version_type, seis_bin_point_version.X_coordinate, seis_bin_point_version.Y_coordinate, seis_bin_point_version.Row_changed_by, seis_bin_point_version.Row_changed_date, seis_bin_point_version.Row_created_by, seis_bin_point_version.Row_created_date, seis_bin_point_version.Row_effective_date, seis_bin_point_version.Row_expiry_date, seis_bin_point_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisBinPointVersion(c *fiber.Ctx) error {
	var seis_bin_point_version dto.Seis_bin_point_version

	setDefaults(&seis_bin_point_version)

	if err := json.Unmarshal(c.Body(), &seis_bin_point_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_bin_point_version.Seis_set_subtype = id

    if seis_bin_point_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_bin_point_version set seis_set_id = :1, bin_grid_id = :2, bin_grid_source = :3, bin_point_id = :4, bin_point_obs_no = :5, acquisition_id = :6, active_ind = :7, easting = :8, easting_ouom = :9, effective_date = :10, elev = :11, elev_ouom = :12, ew_direction = :13, ew_distance = :14, ew_start_line = :15, expiry_date = :16, geog_coord_system_id = :17, local_coord_system_id = :18, map_coord_system_id = :19, northing = :20, northing_ouom = :21, north_type = :22, ns_direction = :23, ns_distance = :24, ns_start_line = :25, polar_azimuth = :26, polar_offset = :27, polar_offset_ouom = :28, ppdm_guid = :29, preferred_ind = :30, reference_loc = :31, remark = :32, seis_point_label = :33, seis_point_lat = :34, seis_point_long = :35, seis_point_no = :36, uwi = :37, version_type = :38, x_coordinate = :39, y_coordinate = :40, row_changed_by = :41, row_changed_date = :42, row_created_by = :43, row_effective_date = :44, row_expiry_date = :45, row_quality = :46 where seis_set_subtype = :48")
	if err != nil {
		return err
	}

	seis_bin_point_version.Row_changed_date = formatDate(seis_bin_point_version.Row_changed_date)
	seis_bin_point_version.Effective_date = formatDateString(seis_bin_point_version.Effective_date)
	seis_bin_point_version.Expiry_date = formatDateString(seis_bin_point_version.Expiry_date)
	seis_bin_point_version.Row_effective_date = formatDateString(seis_bin_point_version.Row_effective_date)
	seis_bin_point_version.Row_expiry_date = formatDateString(seis_bin_point_version.Row_expiry_date)






	rows, err := stmt.Exec(seis_bin_point_version.Seis_set_id, seis_bin_point_version.Bin_grid_id, seis_bin_point_version.Bin_grid_source, seis_bin_point_version.Bin_point_id, seis_bin_point_version.Bin_point_obs_no, seis_bin_point_version.Acquisition_id, seis_bin_point_version.Active_ind, seis_bin_point_version.Easting, seis_bin_point_version.Easting_ouom, seis_bin_point_version.Effective_date, seis_bin_point_version.Elev, seis_bin_point_version.Elev_ouom, seis_bin_point_version.Ew_direction, seis_bin_point_version.Ew_distance, seis_bin_point_version.Ew_start_line, seis_bin_point_version.Expiry_date, seis_bin_point_version.Geog_coord_system_id, seis_bin_point_version.Local_coord_system_id, seis_bin_point_version.Map_coord_system_id, seis_bin_point_version.Northing, seis_bin_point_version.Northing_ouom, seis_bin_point_version.North_type, seis_bin_point_version.Ns_direction, seis_bin_point_version.Ns_distance, seis_bin_point_version.Ns_start_line, seis_bin_point_version.Polar_azimuth, seis_bin_point_version.Polar_offset, seis_bin_point_version.Polar_offset_ouom, seis_bin_point_version.Ppdm_guid, seis_bin_point_version.Preferred_ind, seis_bin_point_version.Reference_loc, seis_bin_point_version.Remark, seis_bin_point_version.Seis_point_label, seis_bin_point_version.Seis_point_lat, seis_bin_point_version.Seis_point_long, seis_bin_point_version.Seis_point_no, seis_bin_point_version.Uwi, seis_bin_point_version.Version_type, seis_bin_point_version.X_coordinate, seis_bin_point_version.Y_coordinate, seis_bin_point_version.Row_changed_by, seis_bin_point_version.Row_changed_date, seis_bin_point_version.Row_created_by, seis_bin_point_version.Row_effective_date, seis_bin_point_version.Row_expiry_date, seis_bin_point_version.Row_quality, seis_bin_point_version.Seis_set_subtype)
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

func PatchSeisBinPointVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_bin_point_version set "
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

func DeleteSeisBinPointVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_bin_point_version dto.Seis_bin_point_version
	seis_bin_point_version.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_bin_point_version where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_bin_point_version.Seis_set_subtype)
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


