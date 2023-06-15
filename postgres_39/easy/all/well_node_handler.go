package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellNode(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_node")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_node

	for rows.Next() {
		var well_node dto.Well_node
		if err := rows.Scan(&well_node.Node_id, &well_node.Active_ind, &well_node.Coordinate_quality, &well_node.Coord_acquisition_id, &well_node.Coord_system_id, &well_node.Effective_date, &well_node.Expiry_date, &well_node.Latitude, &well_node.Legal_survey_type, &well_node.Local_coord_system_id, &well_node.Location_quality, &well_node.Location_type, &well_node.Longitude, &well_node.Node_position, &well_node.Original_obs_no, &well_node.Original_xy_uom, &well_node.Platform_id, &well_node.Platform_sf_subtype, &well_node.Ppdm_guid, &well_node.Preferred_ind, &well_node.Remark, &well_node.Selected_obs_no, &well_node.Source, &well_node.Uwi, &well_node.Row_changed_by, &well_node.Row_changed_date, &well_node.Row_created_by, &well_node.Row_created_date, &well_node.Row_effective_date, &well_node.Row_expiry_date, &well_node.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_node to result
		result = append(result, well_node)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellNode(c *fiber.Ctx) error {
	var well_node dto.Well_node

	setDefaults(&well_node)

	if err := json.Unmarshal(c.Body(), &well_node); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_node values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	well_node.Row_created_date = formatDate(well_node.Row_created_date)
	well_node.Row_changed_date = formatDate(well_node.Row_changed_date)
	well_node.Effective_date = formatDateString(well_node.Effective_date)
	well_node.Expiry_date = formatDateString(well_node.Expiry_date)
	well_node.Row_effective_date = formatDateString(well_node.Row_effective_date)
	well_node.Row_expiry_date = formatDateString(well_node.Row_expiry_date)






	rows, err := stmt.Exec(well_node.Node_id, well_node.Active_ind, well_node.Coordinate_quality, well_node.Coord_acquisition_id, well_node.Coord_system_id, well_node.Effective_date, well_node.Expiry_date, well_node.Latitude, well_node.Legal_survey_type, well_node.Local_coord_system_id, well_node.Location_quality, well_node.Location_type, well_node.Longitude, well_node.Node_position, well_node.Original_obs_no, well_node.Original_xy_uom, well_node.Platform_id, well_node.Platform_sf_subtype, well_node.Ppdm_guid, well_node.Preferred_ind, well_node.Remark, well_node.Selected_obs_no, well_node.Source, well_node.Uwi, well_node.Row_changed_by, well_node.Row_changed_date, well_node.Row_created_by, well_node.Row_created_date, well_node.Row_effective_date, well_node.Row_expiry_date, well_node.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellNode(c *fiber.Ctx) error {
	var well_node dto.Well_node

	setDefaults(&well_node)

	if err := json.Unmarshal(c.Body(), &well_node); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_node.Node_id = id

    if well_node.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_node set active_ind = :1, coordinate_quality = :2, coord_acquisition_id = :3, coord_system_id = :4, effective_date = :5, expiry_date = :6, latitude = :7, legal_survey_type = :8, local_coord_system_id = :9, location_quality = :10, location_type = :11, longitude = :12, node_position = :13, original_obs_no = :14, original_xy_uom = :15, platform_id = :16, platform_sf_subtype = :17, ppdm_guid = :18, preferred_ind = :19, remark = :20, selected_obs_no = :21, source = :22, uwi = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where node_id = :31")
	if err != nil {
		return err
	}

	well_node.Row_changed_date = formatDate(well_node.Row_changed_date)
	well_node.Effective_date = formatDateString(well_node.Effective_date)
	well_node.Expiry_date = formatDateString(well_node.Expiry_date)
	well_node.Row_effective_date = formatDateString(well_node.Row_effective_date)
	well_node.Row_expiry_date = formatDateString(well_node.Row_expiry_date)






	rows, err := stmt.Exec(well_node.Active_ind, well_node.Coordinate_quality, well_node.Coord_acquisition_id, well_node.Coord_system_id, well_node.Effective_date, well_node.Expiry_date, well_node.Latitude, well_node.Legal_survey_type, well_node.Local_coord_system_id, well_node.Location_quality, well_node.Location_type, well_node.Longitude, well_node.Node_position, well_node.Original_obs_no, well_node.Original_xy_uom, well_node.Platform_id, well_node.Platform_sf_subtype, well_node.Ppdm_guid, well_node.Preferred_ind, well_node.Remark, well_node.Selected_obs_no, well_node.Source, well_node.Uwi, well_node.Row_changed_by, well_node.Row_changed_date, well_node.Row_created_by, well_node.Row_effective_date, well_node.Row_expiry_date, well_node.Row_quality, well_node.Node_id)
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

func PatchWellNode(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_node set "
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

func DeleteWellNode(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_node dto.Well_node
	well_node.Node_id = id

	stmt, err := db.Prepare("delete from well_node where node_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_node.Node_id)
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


