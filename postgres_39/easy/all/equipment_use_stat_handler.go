package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipmentUseStat(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment_use_stat")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment_use_stat

	for rows.Next() {
		var equipment_use_stat dto.Equipment_use_stat
		if err := rows.Scan(&equipment_use_stat.Equipment_id, &equipment_use_stat.Spec_id, &equipment_use_stat.Active_ind, &equipment_use_stat.Average_value, &equipment_use_stat.Average_value_ouom, &equipment_use_stat.Average_value_uom, &equipment_use_stat.Cost, &equipment_use_stat.Currency_conversion, &equipment_use_stat.Currency_ouom, &equipment_use_stat.Currency_uom, &equipment_use_stat.Date_format_desc, &equipment_use_stat.Effective_date, &equipment_use_stat.End_date, &equipment_use_stat.End_time, &equipment_use_stat.Equip_use_stat_type, &equipment_use_stat.Expiry_date, &equipment_use_stat.Max_date, &equipment_use_stat.Max_value, &equipment_use_stat.Max_value_ouom, &equipment_use_stat.Max_value_uom, &equipment_use_stat.Min_date, &equipment_use_stat.Min_value, &equipment_use_stat.Min_value_ouom, &equipment_use_stat.Min_value_uom, &equipment_use_stat.Percent_capacity, &equipment_use_stat.Ppdm_guid, &equipment_use_stat.Remark, &equipment_use_stat.Source, &equipment_use_stat.Start_date, &equipment_use_stat.Start_time, &equipment_use_stat.Timezone, &equipment_use_stat.Use_code, &equipment_use_stat.Use_desc, &equipment_use_stat.Row_changed_by, &equipment_use_stat.Row_changed_date, &equipment_use_stat.Row_created_by, &equipment_use_stat.Row_created_date, &equipment_use_stat.Row_effective_date, &equipment_use_stat.Row_expiry_date, &equipment_use_stat.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment_use_stat to result
		result = append(result, equipment_use_stat)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipmentUseStat(c *fiber.Ctx) error {
	var equipment_use_stat dto.Equipment_use_stat

	setDefaults(&equipment_use_stat)

	if err := json.Unmarshal(c.Body(), &equipment_use_stat); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment_use_stat values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40)")
	if err != nil {
		return err
	}
	equipment_use_stat.Row_created_date = formatDate(equipment_use_stat.Row_created_date)
	equipment_use_stat.Row_changed_date = formatDate(equipment_use_stat.Row_changed_date)
	equipment_use_stat.Date_format_desc = formatDateString(equipment_use_stat.Date_format_desc)
	equipment_use_stat.Effective_date = formatDateString(equipment_use_stat.Effective_date)
	equipment_use_stat.End_date = formatDateString(equipment_use_stat.End_date)
	equipment_use_stat.End_time = formatDateString(equipment_use_stat.End_time)
	equipment_use_stat.Expiry_date = formatDateString(equipment_use_stat.Expiry_date)
	equipment_use_stat.Max_date = formatDateString(equipment_use_stat.Max_date)
	equipment_use_stat.Min_date = formatDateString(equipment_use_stat.Min_date)
	equipment_use_stat.Start_date = formatDateString(equipment_use_stat.Start_date)
	equipment_use_stat.Start_time = formatDateString(equipment_use_stat.Start_time)
	equipment_use_stat.Row_effective_date = formatDateString(equipment_use_stat.Row_effective_date)
	equipment_use_stat.Row_expiry_date = formatDateString(equipment_use_stat.Row_expiry_date)






	rows, err := stmt.Exec(equipment_use_stat.Equipment_id, equipment_use_stat.Spec_id, equipment_use_stat.Active_ind, equipment_use_stat.Average_value, equipment_use_stat.Average_value_ouom, equipment_use_stat.Average_value_uom, equipment_use_stat.Cost, equipment_use_stat.Currency_conversion, equipment_use_stat.Currency_ouom, equipment_use_stat.Currency_uom, equipment_use_stat.Date_format_desc, equipment_use_stat.Effective_date, equipment_use_stat.End_date, equipment_use_stat.End_time, equipment_use_stat.Equip_use_stat_type, equipment_use_stat.Expiry_date, equipment_use_stat.Max_date, equipment_use_stat.Max_value, equipment_use_stat.Max_value_ouom, equipment_use_stat.Max_value_uom, equipment_use_stat.Min_date, equipment_use_stat.Min_value, equipment_use_stat.Min_value_ouom, equipment_use_stat.Min_value_uom, equipment_use_stat.Percent_capacity, equipment_use_stat.Ppdm_guid, equipment_use_stat.Remark, equipment_use_stat.Source, equipment_use_stat.Start_date, equipment_use_stat.Start_time, equipment_use_stat.Timezone, equipment_use_stat.Use_code, equipment_use_stat.Use_desc, equipment_use_stat.Row_changed_by, equipment_use_stat.Row_changed_date, equipment_use_stat.Row_created_by, equipment_use_stat.Row_created_date, equipment_use_stat.Row_effective_date, equipment_use_stat.Row_expiry_date, equipment_use_stat.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipmentUseStat(c *fiber.Ctx) error {
	var equipment_use_stat dto.Equipment_use_stat

	setDefaults(&equipment_use_stat)

	if err := json.Unmarshal(c.Body(), &equipment_use_stat); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment_use_stat.Equipment_id = id

    if equipment_use_stat.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment_use_stat set spec_id = :1, active_ind = :2, average_value = :3, average_value_ouom = :4, average_value_uom = :5, cost = :6, currency_conversion = :7, currency_ouom = :8, currency_uom = :9, date_format_desc = :10, effective_date = :11, end_date = :12, end_time = :13, equip_use_stat_type = :14, expiry_date = :15, max_date = :16, max_value = :17, max_value_ouom = :18, max_value_uom = :19, min_date = :20, min_value = :21, min_value_ouom = :22, min_value_uom = :23, percent_capacity = :24, ppdm_guid = :25, remark = :26, source = :27, start_date = :28, start_time = :29, timezone = :30, use_code = :31, use_desc = :32, row_changed_by = :33, row_changed_date = :34, row_created_by = :35, row_effective_date = :36, row_expiry_date = :37, row_quality = :38 where equipment_id = :40")
	if err != nil {
		return err
	}

	equipment_use_stat.Row_changed_date = formatDate(equipment_use_stat.Row_changed_date)
	equipment_use_stat.Date_format_desc = formatDateString(equipment_use_stat.Date_format_desc)
	equipment_use_stat.Effective_date = formatDateString(equipment_use_stat.Effective_date)
	equipment_use_stat.End_date = formatDateString(equipment_use_stat.End_date)
	equipment_use_stat.End_time = formatDateString(equipment_use_stat.End_time)
	equipment_use_stat.Expiry_date = formatDateString(equipment_use_stat.Expiry_date)
	equipment_use_stat.Max_date = formatDateString(equipment_use_stat.Max_date)
	equipment_use_stat.Min_date = formatDateString(equipment_use_stat.Min_date)
	equipment_use_stat.Start_date = formatDateString(equipment_use_stat.Start_date)
	equipment_use_stat.Start_time = formatDateString(equipment_use_stat.Start_time)
	equipment_use_stat.Row_effective_date = formatDateString(equipment_use_stat.Row_effective_date)
	equipment_use_stat.Row_expiry_date = formatDateString(equipment_use_stat.Row_expiry_date)






	rows, err := stmt.Exec(equipment_use_stat.Spec_id, equipment_use_stat.Active_ind, equipment_use_stat.Average_value, equipment_use_stat.Average_value_ouom, equipment_use_stat.Average_value_uom, equipment_use_stat.Cost, equipment_use_stat.Currency_conversion, equipment_use_stat.Currency_ouom, equipment_use_stat.Currency_uom, equipment_use_stat.Date_format_desc, equipment_use_stat.Effective_date, equipment_use_stat.End_date, equipment_use_stat.End_time, equipment_use_stat.Equip_use_stat_type, equipment_use_stat.Expiry_date, equipment_use_stat.Max_date, equipment_use_stat.Max_value, equipment_use_stat.Max_value_ouom, equipment_use_stat.Max_value_uom, equipment_use_stat.Min_date, equipment_use_stat.Min_value, equipment_use_stat.Min_value_ouom, equipment_use_stat.Min_value_uom, equipment_use_stat.Percent_capacity, equipment_use_stat.Ppdm_guid, equipment_use_stat.Remark, equipment_use_stat.Source, equipment_use_stat.Start_date, equipment_use_stat.Start_time, equipment_use_stat.Timezone, equipment_use_stat.Use_code, equipment_use_stat.Use_desc, equipment_use_stat.Row_changed_by, equipment_use_stat.Row_changed_date, equipment_use_stat.Row_created_by, equipment_use_stat.Row_effective_date, equipment_use_stat.Row_expiry_date, equipment_use_stat.Row_quality, equipment_use_stat.Equipment_id)
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

func PatchEquipmentUseStat(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment_use_stat set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "end_date" || key == "end_time" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "start_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where equipment_id = :id"

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

func DeleteEquipmentUseStat(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment_use_stat dto.Equipment_use_stat
	equipment_use_stat.Equipment_id = id

	stmt, err := db.Prepare("delete from equipment_use_stat where equipment_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment_use_stat.Equipment_id)
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


