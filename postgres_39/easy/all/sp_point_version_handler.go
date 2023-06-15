package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpPointVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_point_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_point_version

	for rows.Next() {
		var sp_point_version dto.Sp_point_version
		if err := rows.Scan(&sp_point_version.Point_id, &sp_point_version.Source, &sp_point_version.Version_obs_no, &sp_point_version.Acquisition_id, &sp_point_version.Active_ind, &sp_point_version.Area_id, &sp_point_version.Area_type, &sp_point_version.Easting, &sp_point_version.Easting_ouom, &sp_point_version.Effective_date, &sp_point_version.Elev, &sp_point_version.Elev_ouom, &sp_point_version.Ew_direction, &sp_point_version.Expiry_date, &sp_point_version.Geog_coord_system_id, &sp_point_version.Latitude, &sp_point_version.Legal_survey_type, &sp_point_version.Local_coord_system_id, &sp_point_version.Location_qualifier, &sp_point_version.Location_quality, &sp_point_version.Longitude, &sp_point_version.Map_coord_system_id, &sp_point_version.Md, &sp_point_version.Md_ouom, &sp_point_version.Monument_id, &sp_point_version.Monument_sf_subtype, &sp_point_version.Node_position, &sp_point_version.Northing, &sp_point_version.Northing_ouom, &sp_point_version.North_type, &sp_point_version.Ns_direction, &sp_point_version.Original_ind, &sp_point_version.Original_ref_num, &sp_point_version.Polar_azimuth, &sp_point_version.Polar_offset, &sp_point_version.Polar_offset_ouom, &sp_point_version.Ppdm_guid, &sp_point_version.Preferred_ind, &sp_point_version.Remark, &sp_point_version.Reported_tvd, &sp_point_version.Reported_tvd_ouom, &sp_point_version.Version_type, &sp_point_version.X_coordinate, &sp_point_version.X_offset, &sp_point_version.X_offset_ouom, &sp_point_version.Y_coordinate, &sp_point_version.Y_offset, &sp_point_version.Y_offset_ouom, &sp_point_version.Row_changed_by, &sp_point_version.Row_changed_date, &sp_point_version.Row_created_by, &sp_point_version.Row_created_date, &sp_point_version.Row_effective_date, &sp_point_version.Row_expiry_date, &sp_point_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_point_version to result
		result = append(result, sp_point_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpPointVersion(c *fiber.Ctx) error {
	var sp_point_version dto.Sp_point_version

	setDefaults(&sp_point_version)

	if err := json.Unmarshal(c.Body(), &sp_point_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_point_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55)")
	if err != nil {
		return err
	}
	sp_point_version.Row_created_date = formatDate(sp_point_version.Row_created_date)
	sp_point_version.Row_changed_date = formatDate(sp_point_version.Row_changed_date)
	sp_point_version.Effective_date = formatDateString(sp_point_version.Effective_date)
	sp_point_version.Expiry_date = formatDateString(sp_point_version.Expiry_date)
	sp_point_version.Row_effective_date = formatDateString(sp_point_version.Row_effective_date)
	sp_point_version.Row_expiry_date = formatDateString(sp_point_version.Row_expiry_date)






	rows, err := stmt.Exec(sp_point_version.Point_id, sp_point_version.Source, sp_point_version.Version_obs_no, sp_point_version.Acquisition_id, sp_point_version.Active_ind, sp_point_version.Area_id, sp_point_version.Area_type, sp_point_version.Easting, sp_point_version.Easting_ouom, sp_point_version.Effective_date, sp_point_version.Elev, sp_point_version.Elev_ouom, sp_point_version.Ew_direction, sp_point_version.Expiry_date, sp_point_version.Geog_coord_system_id, sp_point_version.Latitude, sp_point_version.Legal_survey_type, sp_point_version.Local_coord_system_id, sp_point_version.Location_qualifier, sp_point_version.Location_quality, sp_point_version.Longitude, sp_point_version.Map_coord_system_id, sp_point_version.Md, sp_point_version.Md_ouom, sp_point_version.Monument_id, sp_point_version.Monument_sf_subtype, sp_point_version.Node_position, sp_point_version.Northing, sp_point_version.Northing_ouom, sp_point_version.North_type, sp_point_version.Ns_direction, sp_point_version.Original_ind, sp_point_version.Original_ref_num, sp_point_version.Polar_azimuth, sp_point_version.Polar_offset, sp_point_version.Polar_offset_ouom, sp_point_version.Ppdm_guid, sp_point_version.Preferred_ind, sp_point_version.Remark, sp_point_version.Reported_tvd, sp_point_version.Reported_tvd_ouom, sp_point_version.Version_type, sp_point_version.X_coordinate, sp_point_version.X_offset, sp_point_version.X_offset_ouom, sp_point_version.Y_coordinate, sp_point_version.Y_offset, sp_point_version.Y_offset_ouom, sp_point_version.Row_changed_by, sp_point_version.Row_changed_date, sp_point_version.Row_created_by, sp_point_version.Row_created_date, sp_point_version.Row_effective_date, sp_point_version.Row_expiry_date, sp_point_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpPointVersion(c *fiber.Ctx) error {
	var sp_point_version dto.Sp_point_version

	setDefaults(&sp_point_version)

	if err := json.Unmarshal(c.Body(), &sp_point_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_point_version.Point_id = id

    if sp_point_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_point_version set source = :1, version_obs_no = :2, acquisition_id = :3, active_ind = :4, area_id = :5, area_type = :6, easting = :7, easting_ouom = :8, effective_date = :9, elev = :10, elev_ouom = :11, ew_direction = :12, expiry_date = :13, geog_coord_system_id = :14, latitude = :15, legal_survey_type = :16, local_coord_system_id = :17, location_qualifier = :18, location_quality = :19, longitude = :20, map_coord_system_id = :21, md = :22, md_ouom = :23, monument_id = :24, monument_sf_subtype = :25, node_position = :26, northing = :27, northing_ouom = :28, north_type = :29, ns_direction = :30, original_ind = :31, original_ref_num = :32, polar_azimuth = :33, polar_offset = :34, polar_offset_ouom = :35, ppdm_guid = :36, preferred_ind = :37, remark = :38, reported_tvd = :39, reported_tvd_ouom = :40, version_type = :41, x_coordinate = :42, x_offset = :43, x_offset_ouom = :44, y_coordinate = :45, y_offset = :46, y_offset_ouom = :47, row_changed_by = :48, row_changed_date = :49, row_created_by = :50, row_effective_date = :51, row_expiry_date = :52, row_quality = :53 where point_id = :55")
	if err != nil {
		return err
	}

	sp_point_version.Row_changed_date = formatDate(sp_point_version.Row_changed_date)
	sp_point_version.Effective_date = formatDateString(sp_point_version.Effective_date)
	sp_point_version.Expiry_date = formatDateString(sp_point_version.Expiry_date)
	sp_point_version.Row_effective_date = formatDateString(sp_point_version.Row_effective_date)
	sp_point_version.Row_expiry_date = formatDateString(sp_point_version.Row_expiry_date)






	rows, err := stmt.Exec(sp_point_version.Source, sp_point_version.Version_obs_no, sp_point_version.Acquisition_id, sp_point_version.Active_ind, sp_point_version.Area_id, sp_point_version.Area_type, sp_point_version.Easting, sp_point_version.Easting_ouom, sp_point_version.Effective_date, sp_point_version.Elev, sp_point_version.Elev_ouom, sp_point_version.Ew_direction, sp_point_version.Expiry_date, sp_point_version.Geog_coord_system_id, sp_point_version.Latitude, sp_point_version.Legal_survey_type, sp_point_version.Local_coord_system_id, sp_point_version.Location_qualifier, sp_point_version.Location_quality, sp_point_version.Longitude, sp_point_version.Map_coord_system_id, sp_point_version.Md, sp_point_version.Md_ouom, sp_point_version.Monument_id, sp_point_version.Monument_sf_subtype, sp_point_version.Node_position, sp_point_version.Northing, sp_point_version.Northing_ouom, sp_point_version.North_type, sp_point_version.Ns_direction, sp_point_version.Original_ind, sp_point_version.Original_ref_num, sp_point_version.Polar_azimuth, sp_point_version.Polar_offset, sp_point_version.Polar_offset_ouom, sp_point_version.Ppdm_guid, sp_point_version.Preferred_ind, sp_point_version.Remark, sp_point_version.Reported_tvd, sp_point_version.Reported_tvd_ouom, sp_point_version.Version_type, sp_point_version.X_coordinate, sp_point_version.X_offset, sp_point_version.X_offset_ouom, sp_point_version.Y_coordinate, sp_point_version.Y_offset, sp_point_version.Y_offset_ouom, sp_point_version.Row_changed_by, sp_point_version.Row_changed_date, sp_point_version.Row_created_by, sp_point_version.Row_effective_date, sp_point_version.Row_expiry_date, sp_point_version.Row_quality, sp_point_version.Point_id)
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

func PatchSpPointVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_point_version set "
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

func DeleteSpPointVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_point_version dto.Sp_point_version
	sp_point_version.Point_id = id

	stmt, err := db.Prepare("delete from sp_point_version where point_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_point_version.Point_id)
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


