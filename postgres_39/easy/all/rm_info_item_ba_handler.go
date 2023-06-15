package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmInfoItemBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_info_item_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_info_item_ba

	for rows.Next() {
		var rm_info_item_ba dto.Rm_info_item_ba
		if err := rows.Scan(&rm_info_item_ba.Info_item_subtype, &rm_info_item_ba.Information_item_id, &rm_info_item_ba.Contact_id, &rm_info_item_ba.Active_ind, &rm_info_item_ba.Contact_ba_id, &rm_info_item_ba.Contact_ba_type, &rm_info_item_ba.Contact_full_name, &rm_info_item_ba.Contact_type, &rm_info_item_ba.Effective_date, &rm_info_item_ba.Expiry_date, &rm_info_item_ba.First_name, &rm_info_item_ba.Instruction, &rm_info_item_ba.Last_name, &rm_info_item_ba.Middle_name, &rm_info_item_ba.Ppdm_guid, &rm_info_item_ba.Remark, &rm_info_item_ba.Source, &rm_info_item_ba.Row_changed_by, &rm_info_item_ba.Row_changed_date, &rm_info_item_ba.Row_created_by, &rm_info_item_ba.Row_created_date, &rm_info_item_ba.Row_effective_date, &rm_info_item_ba.Row_expiry_date, &rm_info_item_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_info_item_ba to result
		result = append(result, rm_info_item_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmInfoItemBa(c *fiber.Ctx) error {
	var rm_info_item_ba dto.Rm_info_item_ba

	setDefaults(&rm_info_item_ba)

	if err := json.Unmarshal(c.Body(), &rm_info_item_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_info_item_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	rm_info_item_ba.Row_created_date = formatDate(rm_info_item_ba.Row_created_date)
	rm_info_item_ba.Row_changed_date = formatDate(rm_info_item_ba.Row_changed_date)
	rm_info_item_ba.Effective_date = formatDateString(rm_info_item_ba.Effective_date)
	rm_info_item_ba.Expiry_date = formatDateString(rm_info_item_ba.Expiry_date)
	rm_info_item_ba.Row_effective_date = formatDateString(rm_info_item_ba.Row_effective_date)
	rm_info_item_ba.Row_expiry_date = formatDateString(rm_info_item_ba.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_ba.Info_item_subtype, rm_info_item_ba.Information_item_id, rm_info_item_ba.Contact_id, rm_info_item_ba.Active_ind, rm_info_item_ba.Contact_ba_id, rm_info_item_ba.Contact_ba_type, rm_info_item_ba.Contact_full_name, rm_info_item_ba.Contact_type, rm_info_item_ba.Effective_date, rm_info_item_ba.Expiry_date, rm_info_item_ba.First_name, rm_info_item_ba.Instruction, rm_info_item_ba.Last_name, rm_info_item_ba.Middle_name, rm_info_item_ba.Ppdm_guid, rm_info_item_ba.Remark, rm_info_item_ba.Source, rm_info_item_ba.Row_changed_by, rm_info_item_ba.Row_changed_date, rm_info_item_ba.Row_created_by, rm_info_item_ba.Row_created_date, rm_info_item_ba.Row_effective_date, rm_info_item_ba.Row_expiry_date, rm_info_item_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmInfoItemBa(c *fiber.Ctx) error {
	var rm_info_item_ba dto.Rm_info_item_ba

	setDefaults(&rm_info_item_ba)

	if err := json.Unmarshal(c.Body(), &rm_info_item_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_info_item_ba.Info_item_subtype = id

    if rm_info_item_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_info_item_ba set information_item_id = :1, contact_id = :2, active_ind = :3, contact_ba_id = :4, contact_ba_type = :5, contact_full_name = :6, contact_type = :7, effective_date = :8, expiry_date = :9, first_name = :10, instruction = :11, last_name = :12, middle_name = :13, ppdm_guid = :14, remark = :15, source = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where info_item_subtype = :24")
	if err != nil {
		return err
	}

	rm_info_item_ba.Row_changed_date = formatDate(rm_info_item_ba.Row_changed_date)
	rm_info_item_ba.Effective_date = formatDateString(rm_info_item_ba.Effective_date)
	rm_info_item_ba.Expiry_date = formatDateString(rm_info_item_ba.Expiry_date)
	rm_info_item_ba.Row_effective_date = formatDateString(rm_info_item_ba.Row_effective_date)
	rm_info_item_ba.Row_expiry_date = formatDateString(rm_info_item_ba.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_ba.Information_item_id, rm_info_item_ba.Contact_id, rm_info_item_ba.Active_ind, rm_info_item_ba.Contact_ba_id, rm_info_item_ba.Contact_ba_type, rm_info_item_ba.Contact_full_name, rm_info_item_ba.Contact_type, rm_info_item_ba.Effective_date, rm_info_item_ba.Expiry_date, rm_info_item_ba.First_name, rm_info_item_ba.Instruction, rm_info_item_ba.Last_name, rm_info_item_ba.Middle_name, rm_info_item_ba.Ppdm_guid, rm_info_item_ba.Remark, rm_info_item_ba.Source, rm_info_item_ba.Row_changed_by, rm_info_item_ba.Row_changed_date, rm_info_item_ba.Row_created_by, rm_info_item_ba.Row_effective_date, rm_info_item_ba.Row_expiry_date, rm_info_item_ba.Row_quality, rm_info_item_ba.Info_item_subtype)
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

func PatchRmInfoItemBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_info_item_ba set "
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
	query += " where info_item_subtype = :id"

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

func DeleteRmInfoItemBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_info_item_ba dto.Rm_info_item_ba
	rm_info_item_ba.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_info_item_ba where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_info_item_ba.Info_item_subtype)
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


