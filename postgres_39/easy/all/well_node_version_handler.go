package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellNodeVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_node_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_node_version

	for rows.Next() {
		var well_node_version dto.Well_node_version
		if err := rows.Scan(&well_node_version.Node_id, &well_node_version.Source, &well_node_version.Node_obs_no, &well_node_version.Active_ind, &well_node_version.Area_id, &well_node_version.Area_type, &well_node_version.Coordinate_quality, &well_node_version.Coord_acquisition_id, &well_node_version.Easting, &well_node_version.Easting_ouom, &well_node_version.Effective_date, &well_node_version.Elev, &well_node_version.Elev_ouom, &well_node_version.Ew_direction, &well_node_version.Expiry_date, &well_node_version.Geog_coord_system_id, &well_node_version.Latitude, &well_node_version.Legal_survey_type, &well_node_version.Local_coord_system_id, &well_node_version.Location_qualifier, &well_node_version.Location_quality, &well_node_version.Location_type, &well_node_version.Longitude, &well_node_version.Map_coord_system_id, &well_node_version.Md, &well_node_version.Md_ouom, &well_node_version.Monument_id, &well_node_version.Monument_sf_subtype, &well_node_version.Node_position, &well_node_version.Northing, &well_node_version.Northing_ouom, &well_node_version.North_type, &well_node_version.Ns_direction, &well_node_version.Original_xy_uom, &well_node_version.Platform_id, &well_node_version.Platform_sf_subtype, &well_node_version.Polar_azimuth, &well_node_version.Polar_offset, &well_node_version.Polar_offset_ouom, &well_node_version.Ppdm_guid, &well_node_version.Preferred_ind, &well_node_version.Remark, &well_node_version.Reported_tvd, &well_node_version.Reported_tvd_ouom, &well_node_version.Uwi, &well_node_version.Version_type, &well_node_version.X_offset, &well_node_version.X_offset_ouom, &well_node_version.Y_offset, &well_node_version.Y_offset_ouom, &well_node_version.Row_changed_by, &well_node_version.Row_changed_date, &well_node_version.Row_created_by, &well_node_version.Row_created_date, &well_node_version.Row_effective_date, &well_node_version.Row_expiry_date, &well_node_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_node_version to result
		result = append(result, well_node_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellNodeVersion(c *fiber.Ctx) error {
	var well_node_version dto.Well_node_version

	setDefaults(&well_node_version)

	if err := json.Unmarshal(c.Body(), &well_node_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_node_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57)")
	if err != nil {
		return err
	}
	well_node_version.Row_created_date = formatDate(well_node_version.Row_created_date)
	well_node_version.Row_changed_date = formatDate(well_node_version.Row_changed_date)
	well_node_version.Effective_date = formatDateString(well_node_version.Effective_date)
	well_node_version.Expiry_date = formatDateString(well_node_version.Expiry_date)
	well_node_version.Row_effective_date = formatDateString(well_node_version.Row_effective_date)
	well_node_version.Row_expiry_date = formatDateString(well_node_version.Row_expiry_date)






	rows, err := stmt.Exec(well_node_version.Node_id, well_node_version.Source, well_node_version.Node_obs_no, well_node_version.Active_ind, well_node_version.Area_id, well_node_version.Area_type, well_node_version.Coordinate_quality, well_node_version.Coord_acquisition_id, well_node_version.Easting, well_node_version.Easting_ouom, well_node_version.Effective_date, well_node_version.Elev, well_node_version.Elev_ouom, well_node_version.Ew_direction, well_node_version.Expiry_date, well_node_version.Geog_coord_system_id, well_node_version.Latitude, well_node_version.Legal_survey_type, well_node_version.Local_coord_system_id, well_node_version.Location_qualifier, well_node_version.Location_quality, well_node_version.Location_type, well_node_version.Longitude, well_node_version.Map_coord_system_id, well_node_version.Md, well_node_version.Md_ouom, well_node_version.Monument_id, well_node_version.Monument_sf_subtype, well_node_version.Node_position, well_node_version.Northing, well_node_version.Northing_ouom, well_node_version.North_type, well_node_version.Ns_direction, well_node_version.Original_xy_uom, well_node_version.Platform_id, well_node_version.Platform_sf_subtype, well_node_version.Polar_azimuth, well_node_version.Polar_offset, well_node_version.Polar_offset_ouom, well_node_version.Ppdm_guid, well_node_version.Preferred_ind, well_node_version.Remark, well_node_version.Reported_tvd, well_node_version.Reported_tvd_ouom, well_node_version.Uwi, well_node_version.Version_type, well_node_version.X_offset, well_node_version.X_offset_ouom, well_node_version.Y_offset, well_node_version.Y_offset_ouom, well_node_version.Row_changed_by, well_node_version.Row_changed_date, well_node_version.Row_created_by, well_node_version.Row_created_date, well_node_version.Row_effective_date, well_node_version.Row_expiry_date, well_node_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellNodeVersion(c *fiber.Ctx) error {
	var well_node_version dto.Well_node_version

	setDefaults(&well_node_version)

	if err := json.Unmarshal(c.Body(), &well_node_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_node_version.Node_id = id

    if well_node_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_node_version set source = :1, node_obs_no = :2, active_ind = :3, area_id = :4, area_type = :5, coordinate_quality = :6, coord_acquisition_id = :7, easting = :8, easting_ouom = :9, effective_date = :10, elev = :11, elev_ouom = :12, ew_direction = :13, expiry_date = :14, geog_coord_system_id = :15, latitude = :16, legal_survey_type = :17, local_coord_system_id = :18, location_qualifier = :19, location_quality = :20, location_type = :21, longitude = :22, map_coord_system_id = :23, md = :24, md_ouom = :25, monument_id = :26, monument_sf_subtype = :27, node_position = :28, northing = :29, northing_ouom = :30, north_type = :31, ns_direction = :32, original_xy_uom = :33, platform_id = :34, platform_sf_subtype = :35, polar_azimuth = :36, polar_offset = :37, polar_offset_ouom = :38, ppdm_guid = :39, preferred_ind = :40, remark = :41, reported_tvd = :42, reported_tvd_ouom = :43, uwi = :44, version_type = :45, x_offset = :46, x_offset_ouom = :47, y_offset = :48, y_offset_ouom = :49, row_changed_by = :50, row_changed_date = :51, row_created_by = :52, row_effective_date = :53, row_expiry_date = :54, row_quality = :55 where node_id = :57")
	if err != nil {
		return err
	}

	well_node_version.Row_changed_date = formatDate(well_node_version.Row_changed_date)
	well_node_version.Effective_date = formatDateString(well_node_version.Effective_date)
	well_node_version.Expiry_date = formatDateString(well_node_version.Expiry_date)
	well_node_version.Row_effective_date = formatDateString(well_node_version.Row_effective_date)
	well_node_version.Row_expiry_date = formatDateString(well_node_version.Row_expiry_date)






	rows, err := stmt.Exec(well_node_version.Source, well_node_version.Node_obs_no, well_node_version.Active_ind, well_node_version.Area_id, well_node_version.Area_type, well_node_version.Coordinate_quality, well_node_version.Coord_acquisition_id, well_node_version.Easting, well_node_version.Easting_ouom, well_node_version.Effective_date, well_node_version.Elev, well_node_version.Elev_ouom, well_node_version.Ew_direction, well_node_version.Expiry_date, well_node_version.Geog_coord_system_id, well_node_version.Latitude, well_node_version.Legal_survey_type, well_node_version.Local_coord_system_id, well_node_version.Location_qualifier, well_node_version.Location_quality, well_node_version.Location_type, well_node_version.Longitude, well_node_version.Map_coord_system_id, well_node_version.Md, well_node_version.Md_ouom, well_node_version.Monument_id, well_node_version.Monument_sf_subtype, well_node_version.Node_position, well_node_version.Northing, well_node_version.Northing_ouom, well_node_version.North_type, well_node_version.Ns_direction, well_node_version.Original_xy_uom, well_node_version.Platform_id, well_node_version.Platform_sf_subtype, well_node_version.Polar_azimuth, well_node_version.Polar_offset, well_node_version.Polar_offset_ouom, well_node_version.Ppdm_guid, well_node_version.Preferred_ind, well_node_version.Remark, well_node_version.Reported_tvd, well_node_version.Reported_tvd_ouom, well_node_version.Uwi, well_node_version.Version_type, well_node_version.X_offset, well_node_version.X_offset_ouom, well_node_version.Y_offset, well_node_version.Y_offset_ouom, well_node_version.Row_changed_by, well_node_version.Row_changed_date, well_node_version.Row_created_by, well_node_version.Row_effective_date, well_node_version.Row_expiry_date, well_node_version.Row_quality, well_node_version.Node_id)
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

func PatchWellNodeVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_node_version set "
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
	query += " where node_id = :id"

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

func DeleteWellNodeVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_node_version dto.Well_node_version
	well_node_version.Node_id = id

	stmt, err := db.Prepare("delete from well_node_version where node_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_node_version.Node_id)
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


