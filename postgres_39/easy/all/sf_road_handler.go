package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfRoad(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_road")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_road

	for rows.Next() {
		var sf_road dto.Sf_road
		if err := rows.Scan(&sf_road.Sf_subtype, &sf_road.Road_id, &sf_road.Active_ind, &sf_road.Capacity, &sf_road.Capacity_ouom, &sf_road.Communication_freq, &sf_road.Communication_freq_desc, &sf_road.Control_type, &sf_road.Current_operator_ba_id, &sf_road.Direction, &sf_road.Driving_side, &sf_road.Effective_date, &sf_road.Expiry_date, &sf_road.Ppdm_guid, &sf_road.Remark, &sf_road.Road_length, &sf_road.Road_length_ouom, &sf_road.Road_type, &sf_road.Road_width, &sf_road.Road_width_ouom, &sf_road.Source, &sf_road.Surface_type, &sf_road.Traffic_flow_type, &sf_road.Row_changed_by, &sf_road.Row_changed_date, &sf_road.Row_created_by, &sf_road.Row_created_date, &sf_road.Row_effective_date, &sf_road.Row_expiry_date, &sf_road.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_road to result
		result = append(result, sf_road)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfRoad(c *fiber.Ctx) error {
	var sf_road dto.Sf_road

	setDefaults(&sf_road)

	if err := json.Unmarshal(c.Body(), &sf_road); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_road values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	sf_road.Row_created_date = formatDate(sf_road.Row_created_date)
	sf_road.Row_changed_date = formatDate(sf_road.Row_changed_date)
	sf_road.Effective_date = formatDateString(sf_road.Effective_date)
	sf_road.Expiry_date = formatDateString(sf_road.Expiry_date)
	sf_road.Row_effective_date = formatDateString(sf_road.Row_effective_date)
	sf_road.Row_expiry_date = formatDateString(sf_road.Row_expiry_date)






	rows, err := stmt.Exec(sf_road.Sf_subtype, sf_road.Road_id, sf_road.Active_ind, sf_road.Capacity, sf_road.Capacity_ouom, sf_road.Communication_freq, sf_road.Communication_freq_desc, sf_road.Control_type, sf_road.Current_operator_ba_id, sf_road.Direction, sf_road.Driving_side, sf_road.Effective_date, sf_road.Expiry_date, sf_road.Ppdm_guid, sf_road.Remark, sf_road.Road_length, sf_road.Road_length_ouom, sf_road.Road_type, sf_road.Road_width, sf_road.Road_width_ouom, sf_road.Source, sf_road.Surface_type, sf_road.Traffic_flow_type, sf_road.Row_changed_by, sf_road.Row_changed_date, sf_road.Row_created_by, sf_road.Row_created_date, sf_road.Row_effective_date, sf_road.Row_expiry_date, sf_road.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfRoad(c *fiber.Ctx) error {
	var sf_road dto.Sf_road

	setDefaults(&sf_road)

	if err := json.Unmarshal(c.Body(), &sf_road); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_road.Sf_subtype = id

    if sf_road.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_road set road_id = :1, active_ind = :2, capacity = :3, capacity_ouom = :4, communication_freq = :5, communication_freq_desc = :6, control_type = :7, current_operator_ba_id = :8, direction = :9, driving_side = :10, effective_date = :11, expiry_date = :12, ppdm_guid = :13, remark = :14, road_length = :15, road_length_ouom = :16, road_type = :17, road_width = :18, road_width_ouom = :19, source = :20, surface_type = :21, traffic_flow_type = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where sf_subtype = :30")
	if err != nil {
		return err
	}

	sf_road.Row_changed_date = formatDate(sf_road.Row_changed_date)
	sf_road.Effective_date = formatDateString(sf_road.Effective_date)
	sf_road.Expiry_date = formatDateString(sf_road.Expiry_date)
	sf_road.Row_effective_date = formatDateString(sf_road.Row_effective_date)
	sf_road.Row_expiry_date = formatDateString(sf_road.Row_expiry_date)






	rows, err := stmt.Exec(sf_road.Road_id, sf_road.Active_ind, sf_road.Capacity, sf_road.Capacity_ouom, sf_road.Communication_freq, sf_road.Communication_freq_desc, sf_road.Control_type, sf_road.Current_operator_ba_id, sf_road.Direction, sf_road.Driving_side, sf_road.Effective_date, sf_road.Expiry_date, sf_road.Ppdm_guid, sf_road.Remark, sf_road.Road_length, sf_road.Road_length_ouom, sf_road.Road_type, sf_road.Road_width, sf_road.Road_width_ouom, sf_road.Source, sf_road.Surface_type, sf_road.Traffic_flow_type, sf_road.Row_changed_by, sf_road.Row_changed_date, sf_road.Row_created_by, sf_road.Row_effective_date, sf_road.Row_expiry_date, sf_road.Row_quality, sf_road.Sf_subtype)
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

func PatchSfRoad(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_road set "
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
	query += " where sf_subtype = :id"

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

func DeleteSfRoad(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_road dto.Sf_road
	sf_road.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_road where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_road.Sf_subtype)
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


