package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratFieldNode(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_field_node")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_field_node

	for rows.Next() {
		var strat_field_node dto.Strat_field_node
		if err := rows.Scan(&strat_field_node.Field_station_id, &strat_field_node.Node_id, &strat_field_node.Source, &strat_field_node.Active_ind, &strat_field_node.Coord_acquisition_id, &strat_field_node.Coord_system_id, &strat_field_node.Effective_date, &strat_field_node.Elev, &strat_field_node.Elev_ouom, &strat_field_node.Expiry_date, &strat_field_node.Latitude, &strat_field_node.Local_coord_system_id, &strat_field_node.Longitude, &strat_field_node.Node_loc_type, &strat_field_node.Original_obs_no, &strat_field_node.Original_xy_uom, &strat_field_node.Ppdm_guid, &strat_field_node.Remark, &strat_field_node.Selected_obs_no, &strat_field_node.Row_changed_by, &strat_field_node.Row_changed_date, &strat_field_node.Row_created_by, &strat_field_node.Row_created_date, &strat_field_node.Row_effective_date, &strat_field_node.Row_expiry_date, &strat_field_node.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_field_node to result
		result = append(result, strat_field_node)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratFieldNode(c *fiber.Ctx) error {
	var strat_field_node dto.Strat_field_node

	setDefaults(&strat_field_node)

	if err := json.Unmarshal(c.Body(), &strat_field_node); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_field_node values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	strat_field_node.Row_created_date = formatDate(strat_field_node.Row_created_date)
	strat_field_node.Row_changed_date = formatDate(strat_field_node.Row_changed_date)
	strat_field_node.Effective_date = formatDateString(strat_field_node.Effective_date)
	strat_field_node.Expiry_date = formatDateString(strat_field_node.Expiry_date)
	strat_field_node.Row_effective_date = formatDateString(strat_field_node.Row_effective_date)
	strat_field_node.Row_expiry_date = formatDateString(strat_field_node.Row_expiry_date)






	rows, err := stmt.Exec(strat_field_node.Field_station_id, strat_field_node.Node_id, strat_field_node.Source, strat_field_node.Active_ind, strat_field_node.Coord_acquisition_id, strat_field_node.Coord_system_id, strat_field_node.Effective_date, strat_field_node.Elev, strat_field_node.Elev_ouom, strat_field_node.Expiry_date, strat_field_node.Latitude, strat_field_node.Local_coord_system_id, strat_field_node.Longitude, strat_field_node.Node_loc_type, strat_field_node.Original_obs_no, strat_field_node.Original_xy_uom, strat_field_node.Ppdm_guid, strat_field_node.Remark, strat_field_node.Selected_obs_no, strat_field_node.Row_changed_by, strat_field_node.Row_changed_date, strat_field_node.Row_created_by, strat_field_node.Row_created_date, strat_field_node.Row_effective_date, strat_field_node.Row_expiry_date, strat_field_node.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratFieldNode(c *fiber.Ctx) error {
	var strat_field_node dto.Strat_field_node

	setDefaults(&strat_field_node)

	if err := json.Unmarshal(c.Body(), &strat_field_node); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_field_node.Field_station_id = id

    if strat_field_node.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_field_node set node_id = :1, source = :2, active_ind = :3, coord_acquisition_id = :4, coord_system_id = :5, effective_date = :6, elev = :7, elev_ouom = :8, expiry_date = :9, latitude = :10, local_coord_system_id = :11, longitude = :12, node_loc_type = :13, original_obs_no = :14, original_xy_uom = :15, ppdm_guid = :16, remark = :17, selected_obs_no = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where field_station_id = :26")
	if err != nil {
		return err
	}

	strat_field_node.Row_changed_date = formatDate(strat_field_node.Row_changed_date)
	strat_field_node.Effective_date = formatDateString(strat_field_node.Effective_date)
	strat_field_node.Expiry_date = formatDateString(strat_field_node.Expiry_date)
	strat_field_node.Row_effective_date = formatDateString(strat_field_node.Row_effective_date)
	strat_field_node.Row_expiry_date = formatDateString(strat_field_node.Row_expiry_date)






	rows, err := stmt.Exec(strat_field_node.Node_id, strat_field_node.Source, strat_field_node.Active_ind, strat_field_node.Coord_acquisition_id, strat_field_node.Coord_system_id, strat_field_node.Effective_date, strat_field_node.Elev, strat_field_node.Elev_ouom, strat_field_node.Expiry_date, strat_field_node.Latitude, strat_field_node.Local_coord_system_id, strat_field_node.Longitude, strat_field_node.Node_loc_type, strat_field_node.Original_obs_no, strat_field_node.Original_xy_uom, strat_field_node.Ppdm_guid, strat_field_node.Remark, strat_field_node.Selected_obs_no, strat_field_node.Row_changed_by, strat_field_node.Row_changed_date, strat_field_node.Row_created_by, strat_field_node.Row_effective_date, strat_field_node.Row_expiry_date, strat_field_node.Row_quality, strat_field_node.Field_station_id)
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

func PatchStratFieldNode(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_field_node set "
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

func DeleteStratFieldNode(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_field_node dto.Strat_field_node
	strat_field_node.Field_station_id = id

	stmt, err := db.Prepare("delete from strat_field_node where field_station_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_field_node.Field_station_id)
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


