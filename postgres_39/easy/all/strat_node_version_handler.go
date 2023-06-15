package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratNodeVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_node_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_node_version

	for rows.Next() {
		var strat_node_version dto.Strat_node_version
		if err := rows.Scan(&strat_node_version.Field_station_id, &strat_node_version.Node_id, &strat_node_version.Source, &strat_node_version.Node_obs_no, &strat_node_version.Acquisition_id, &strat_node_version.Active_ind, &strat_node_version.Easting, &strat_node_version.Easting_ouom, &strat_node_version.Effective_date, &strat_node_version.Elev, &strat_node_version.Elev_ouom, &strat_node_version.Ew_direction, &strat_node_version.Expiry_date, &strat_node_version.Geog_coord_system_id, &strat_node_version.Latitude, &strat_node_version.Legal_survey_type, &strat_node_version.Local_coord_system_id, &strat_node_version.Location_qualifier, &strat_node_version.Longitude, &strat_node_version.Map_coord_system_id, &strat_node_version.Md, &strat_node_version.Md_ouom, &strat_node_version.Node_loc_type, &strat_node_version.Northing, &strat_node_version.Northing_ouom, &strat_node_version.North_type, &strat_node_version.Ns_direction, &strat_node_version.Polar_azimuth, &strat_node_version.Polar_offset, &strat_node_version.Polar_offset_ouom, &strat_node_version.Ppdm_guid, &strat_node_version.Remark, &strat_node_version.Reported_tvd, &strat_node_version.Reported_tvd_ouom, &strat_node_version.Version_type, &strat_node_version.X_offset, &strat_node_version.X_offset_ouom, &strat_node_version.Y_offset, &strat_node_version.Y_offset_ouom, &strat_node_version.Row_changed_by, &strat_node_version.Row_changed_date, &strat_node_version.Row_created_by, &strat_node_version.Row_created_date, &strat_node_version.Row_effective_date, &strat_node_version.Row_expiry_date, &strat_node_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_node_version to result
		result = append(result, strat_node_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratNodeVersion(c *fiber.Ctx) error {
	var strat_node_version dto.Strat_node_version

	setDefaults(&strat_node_version)

	if err := json.Unmarshal(c.Body(), &strat_node_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_node_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46)")
	if err != nil {
		return err
	}
	strat_node_version.Row_created_date = formatDate(strat_node_version.Row_created_date)
	strat_node_version.Row_changed_date = formatDate(strat_node_version.Row_changed_date)
	strat_node_version.Effective_date = formatDateString(strat_node_version.Effective_date)
	strat_node_version.Expiry_date = formatDateString(strat_node_version.Expiry_date)
	strat_node_version.Row_effective_date = formatDateString(strat_node_version.Row_effective_date)
	strat_node_version.Row_expiry_date = formatDateString(strat_node_version.Row_expiry_date)






	rows, err := stmt.Exec(strat_node_version.Field_station_id, strat_node_version.Node_id, strat_node_version.Source, strat_node_version.Node_obs_no, strat_node_version.Acquisition_id, strat_node_version.Active_ind, strat_node_version.Easting, strat_node_version.Easting_ouom, strat_node_version.Effective_date, strat_node_version.Elev, strat_node_version.Elev_ouom, strat_node_version.Ew_direction, strat_node_version.Expiry_date, strat_node_version.Geog_coord_system_id, strat_node_version.Latitude, strat_node_version.Legal_survey_type, strat_node_version.Local_coord_system_id, strat_node_version.Location_qualifier, strat_node_version.Longitude, strat_node_version.Map_coord_system_id, strat_node_version.Md, strat_node_version.Md_ouom, strat_node_version.Node_loc_type, strat_node_version.Northing, strat_node_version.Northing_ouom, strat_node_version.North_type, strat_node_version.Ns_direction, strat_node_version.Polar_azimuth, strat_node_version.Polar_offset, strat_node_version.Polar_offset_ouom, strat_node_version.Ppdm_guid, strat_node_version.Remark, strat_node_version.Reported_tvd, strat_node_version.Reported_tvd_ouom, strat_node_version.Version_type, strat_node_version.X_offset, strat_node_version.X_offset_ouom, strat_node_version.Y_offset, strat_node_version.Y_offset_ouom, strat_node_version.Row_changed_by, strat_node_version.Row_changed_date, strat_node_version.Row_created_by, strat_node_version.Row_created_date, strat_node_version.Row_effective_date, strat_node_version.Row_expiry_date, strat_node_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratNodeVersion(c *fiber.Ctx) error {
	var strat_node_version dto.Strat_node_version

	setDefaults(&strat_node_version)

	if err := json.Unmarshal(c.Body(), &strat_node_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_node_version.Field_station_id = id

    if strat_node_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_node_version set node_id = :1, source = :2, node_obs_no = :3, acquisition_id = :4, active_ind = :5, easting = :6, easting_ouom = :7, effective_date = :8, elev = :9, elev_ouom = :10, ew_direction = :11, expiry_date = :12, geog_coord_system_id = :13, latitude = :14, legal_survey_type = :15, local_coord_system_id = :16, location_qualifier = :17, longitude = :18, map_coord_system_id = :19, md = :20, md_ouom = :21, node_loc_type = :22, northing = :23, northing_ouom = :24, north_type = :25, ns_direction = :26, polar_azimuth = :27, polar_offset = :28, polar_offset_ouom = :29, ppdm_guid = :30, remark = :31, reported_tvd = :32, reported_tvd_ouom = :33, version_type = :34, x_offset = :35, x_offset_ouom = :36, y_offset = :37, y_offset_ouom = :38, row_changed_by = :39, row_changed_date = :40, row_created_by = :41, row_effective_date = :42, row_expiry_date = :43, row_quality = :44 where field_station_id = :46")
	if err != nil {
		return err
	}

	strat_node_version.Row_changed_date = formatDate(strat_node_version.Row_changed_date)
	strat_node_version.Effective_date = formatDateString(strat_node_version.Effective_date)
	strat_node_version.Expiry_date = formatDateString(strat_node_version.Expiry_date)
	strat_node_version.Row_effective_date = formatDateString(strat_node_version.Row_effective_date)
	strat_node_version.Row_expiry_date = formatDateString(strat_node_version.Row_expiry_date)






	rows, err := stmt.Exec(strat_node_version.Node_id, strat_node_version.Source, strat_node_version.Node_obs_no, strat_node_version.Acquisition_id, strat_node_version.Active_ind, strat_node_version.Easting, strat_node_version.Easting_ouom, strat_node_version.Effective_date, strat_node_version.Elev, strat_node_version.Elev_ouom, strat_node_version.Ew_direction, strat_node_version.Expiry_date, strat_node_version.Geog_coord_system_id, strat_node_version.Latitude, strat_node_version.Legal_survey_type, strat_node_version.Local_coord_system_id, strat_node_version.Location_qualifier, strat_node_version.Longitude, strat_node_version.Map_coord_system_id, strat_node_version.Md, strat_node_version.Md_ouom, strat_node_version.Node_loc_type, strat_node_version.Northing, strat_node_version.Northing_ouom, strat_node_version.North_type, strat_node_version.Ns_direction, strat_node_version.Polar_azimuth, strat_node_version.Polar_offset, strat_node_version.Polar_offset_ouom, strat_node_version.Ppdm_guid, strat_node_version.Remark, strat_node_version.Reported_tvd, strat_node_version.Reported_tvd_ouom, strat_node_version.Version_type, strat_node_version.X_offset, strat_node_version.X_offset_ouom, strat_node_version.Y_offset, strat_node_version.Y_offset_ouom, strat_node_version.Row_changed_by, strat_node_version.Row_changed_date, strat_node_version.Row_created_by, strat_node_version.Row_effective_date, strat_node_version.Row_expiry_date, strat_node_version.Row_quality, strat_node_version.Field_station_id)
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

func PatchStratNodeVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_node_version set "
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
	query += " where field_station_id = :id"

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

func DeleteStratNodeVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_node_version dto.Strat_node_version
	strat_node_version.Field_station_id = id

	stmt, err := db.Prepare("delete from strat_node_version where field_station_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_node_version.Field_station_id)
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


