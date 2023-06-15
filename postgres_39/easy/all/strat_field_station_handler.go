package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratFieldStation(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_field_station")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_field_station

	for rows.Next() {
		var strat_field_station dto.Strat_field_station
		if err := rows.Scan(&strat_field_station.Field_station_id, &strat_field_station.Active_ind, &strat_field_station.Air_photo_num, &strat_field_station.Area_id, &strat_field_station.Area_type, &strat_field_station.Complete_date, &strat_field_station.Description, &strat_field_station.Effective_date, &strat_field_station.Expiry_date, &strat_field_station.Field_station_type, &strat_field_station.Ppdm_guid, &strat_field_station.Remark, &strat_field_station.Source, &strat_field_station.Source_document_id, &strat_field_station.Start_date, &strat_field_station.Row_changed_by, &strat_field_station.Row_changed_date, &strat_field_station.Row_created_by, &strat_field_station.Row_created_date, &strat_field_station.Row_effective_date, &strat_field_station.Row_expiry_date, &strat_field_station.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_field_station to result
		result = append(result, strat_field_station)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratFieldStation(c *fiber.Ctx) error {
	var strat_field_station dto.Strat_field_station

	setDefaults(&strat_field_station)

	if err := json.Unmarshal(c.Body(), &strat_field_station); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_field_station values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	strat_field_station.Row_created_date = formatDate(strat_field_station.Row_created_date)
	strat_field_station.Row_changed_date = formatDate(strat_field_station.Row_changed_date)
	strat_field_station.Complete_date = formatDateString(strat_field_station.Complete_date)
	strat_field_station.Effective_date = formatDateString(strat_field_station.Effective_date)
	strat_field_station.Expiry_date = formatDateString(strat_field_station.Expiry_date)
	strat_field_station.Start_date = formatDateString(strat_field_station.Start_date)
	strat_field_station.Row_effective_date = formatDateString(strat_field_station.Row_effective_date)
	strat_field_station.Row_expiry_date = formatDateString(strat_field_station.Row_expiry_date)






	rows, err := stmt.Exec(strat_field_station.Field_station_id, strat_field_station.Active_ind, strat_field_station.Air_photo_num, strat_field_station.Area_id, strat_field_station.Area_type, strat_field_station.Complete_date, strat_field_station.Description, strat_field_station.Effective_date, strat_field_station.Expiry_date, strat_field_station.Field_station_type, strat_field_station.Ppdm_guid, strat_field_station.Remark, strat_field_station.Source, strat_field_station.Source_document_id, strat_field_station.Start_date, strat_field_station.Row_changed_by, strat_field_station.Row_changed_date, strat_field_station.Row_created_by, strat_field_station.Row_created_date, strat_field_station.Row_effective_date, strat_field_station.Row_expiry_date, strat_field_station.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratFieldStation(c *fiber.Ctx) error {
	var strat_field_station dto.Strat_field_station

	setDefaults(&strat_field_station)

	if err := json.Unmarshal(c.Body(), &strat_field_station); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_field_station.Field_station_id = id

    if strat_field_station.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_field_station set active_ind = :1, air_photo_num = :2, area_id = :3, area_type = :4, complete_date = :5, description = :6, effective_date = :7, expiry_date = :8, field_station_type = :9, ppdm_guid = :10, remark = :11, source = :12, source_document_id = :13, start_date = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where field_station_id = :22")
	if err != nil {
		return err
	}

	strat_field_station.Row_changed_date = formatDate(strat_field_station.Row_changed_date)
	strat_field_station.Complete_date = formatDateString(strat_field_station.Complete_date)
	strat_field_station.Effective_date = formatDateString(strat_field_station.Effective_date)
	strat_field_station.Expiry_date = formatDateString(strat_field_station.Expiry_date)
	strat_field_station.Start_date = formatDateString(strat_field_station.Start_date)
	strat_field_station.Row_effective_date = formatDateString(strat_field_station.Row_effective_date)
	strat_field_station.Row_expiry_date = formatDateString(strat_field_station.Row_expiry_date)






	rows, err := stmt.Exec(strat_field_station.Active_ind, strat_field_station.Air_photo_num, strat_field_station.Area_id, strat_field_station.Area_type, strat_field_station.Complete_date, strat_field_station.Description, strat_field_station.Effective_date, strat_field_station.Expiry_date, strat_field_station.Field_station_type, strat_field_station.Ppdm_guid, strat_field_station.Remark, strat_field_station.Source, strat_field_station.Source_document_id, strat_field_station.Start_date, strat_field_station.Row_changed_by, strat_field_station.Row_changed_date, strat_field_station.Row_created_by, strat_field_station.Row_effective_date, strat_field_station.Row_expiry_date, strat_field_station.Row_quality, strat_field_station.Field_station_id)
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

func PatchStratFieldStation(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_field_station set "
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
		} else if key == "complete_date" || key == "effective_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteStratFieldStation(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_field_station dto.Strat_field_station
	strat_field_station.Field_station_id = id

	stmt, err := db.Prepare("delete from strat_field_station where field_station_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_field_station.Field_station_id)
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


