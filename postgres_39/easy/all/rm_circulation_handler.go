package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmCirculation(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_circulation")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_circulation

	for rows.Next() {
		var rm_circulation dto.Rm_circulation
		if err := rows.Scan(&rm_circulation.Circ_id, &rm_circulation.Active_ind, &rm_circulation.Authorized_by, &rm_circulation.Checked_out_by, &rm_circulation.Circ_due_date, &rm_circulation.Circ_in_date, &rm_circulation.Circ_out_date, &rm_circulation.Condition_in, &rm_circulation.Condition_out, &rm_circulation.Data_circ_status, &rm_circulation.Data_content_seq_no, &rm_circulation.Effective_date, &rm_circulation.Expiry_date, &rm_circulation.Information_item_id, &rm_circulation.Info_item_subtype, &rm_circulation.Physical_item_id, &rm_circulation.Ppdm_guid, &rm_circulation.Project_id, &rm_circulation.Project_step_id, &rm_circulation.Reference_num, &rm_circulation.Remark, &rm_circulation.Reserved_by, &rm_circulation.Reserved_ind, &rm_circulation.Source, &rm_circulation.Status_date, &rm_circulation.Row_changed_by, &rm_circulation.Row_changed_date, &rm_circulation.Row_created_by, &rm_circulation.Row_created_date, &rm_circulation.Row_effective_date, &rm_circulation.Row_expiry_date, &rm_circulation.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_circulation to result
		result = append(result, rm_circulation)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmCirculation(c *fiber.Ctx) error {
	var rm_circulation dto.Rm_circulation

	setDefaults(&rm_circulation)

	if err := json.Unmarshal(c.Body(), &rm_circulation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_circulation values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	rm_circulation.Row_created_date = formatDate(rm_circulation.Row_created_date)
	rm_circulation.Row_changed_date = formatDate(rm_circulation.Row_changed_date)
	rm_circulation.Circ_due_date = formatDateString(rm_circulation.Circ_due_date)
	rm_circulation.Circ_in_date = formatDateString(rm_circulation.Circ_in_date)
	rm_circulation.Circ_out_date = formatDateString(rm_circulation.Circ_out_date)
	rm_circulation.Effective_date = formatDateString(rm_circulation.Effective_date)
	rm_circulation.Expiry_date = formatDateString(rm_circulation.Expiry_date)
	rm_circulation.Status_date = formatDateString(rm_circulation.Status_date)
	rm_circulation.Row_effective_date = formatDateString(rm_circulation.Row_effective_date)
	rm_circulation.Row_expiry_date = formatDateString(rm_circulation.Row_expiry_date)






	rows, err := stmt.Exec(rm_circulation.Circ_id, rm_circulation.Active_ind, rm_circulation.Authorized_by, rm_circulation.Checked_out_by, rm_circulation.Circ_due_date, rm_circulation.Circ_in_date, rm_circulation.Circ_out_date, rm_circulation.Condition_in, rm_circulation.Condition_out, rm_circulation.Data_circ_status, rm_circulation.Data_content_seq_no, rm_circulation.Effective_date, rm_circulation.Expiry_date, rm_circulation.Information_item_id, rm_circulation.Info_item_subtype, rm_circulation.Physical_item_id, rm_circulation.Ppdm_guid, rm_circulation.Project_id, rm_circulation.Project_step_id, rm_circulation.Reference_num, rm_circulation.Remark, rm_circulation.Reserved_by, rm_circulation.Reserved_ind, rm_circulation.Source, rm_circulation.Status_date, rm_circulation.Row_changed_by, rm_circulation.Row_changed_date, rm_circulation.Row_created_by, rm_circulation.Row_created_date, rm_circulation.Row_effective_date, rm_circulation.Row_expiry_date, rm_circulation.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmCirculation(c *fiber.Ctx) error {
	var rm_circulation dto.Rm_circulation

	setDefaults(&rm_circulation)

	if err := json.Unmarshal(c.Body(), &rm_circulation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_circulation.Circ_id = id

    if rm_circulation.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_circulation set active_ind = :1, authorized_by = :2, checked_out_by = :3, circ_due_date = :4, circ_in_date = :5, circ_out_date = :6, condition_in = :7, condition_out = :8, data_circ_status = :9, data_content_seq_no = :10, effective_date = :11, expiry_date = :12, information_item_id = :13, info_item_subtype = :14, physical_item_id = :15, ppdm_guid = :16, project_id = :17, project_step_id = :18, reference_num = :19, remark = :20, reserved_by = :21, reserved_ind = :22, source = :23, status_date = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where circ_id = :32")
	if err != nil {
		return err
	}

	rm_circulation.Row_changed_date = formatDate(rm_circulation.Row_changed_date)
	rm_circulation.Circ_due_date = formatDateString(rm_circulation.Circ_due_date)
	rm_circulation.Circ_in_date = formatDateString(rm_circulation.Circ_in_date)
	rm_circulation.Circ_out_date = formatDateString(rm_circulation.Circ_out_date)
	rm_circulation.Effective_date = formatDateString(rm_circulation.Effective_date)
	rm_circulation.Expiry_date = formatDateString(rm_circulation.Expiry_date)
	rm_circulation.Status_date = formatDateString(rm_circulation.Status_date)
	rm_circulation.Row_effective_date = formatDateString(rm_circulation.Row_effective_date)
	rm_circulation.Row_expiry_date = formatDateString(rm_circulation.Row_expiry_date)






	rows, err := stmt.Exec(rm_circulation.Active_ind, rm_circulation.Authorized_by, rm_circulation.Checked_out_by, rm_circulation.Circ_due_date, rm_circulation.Circ_in_date, rm_circulation.Circ_out_date, rm_circulation.Condition_in, rm_circulation.Condition_out, rm_circulation.Data_circ_status, rm_circulation.Data_content_seq_no, rm_circulation.Effective_date, rm_circulation.Expiry_date, rm_circulation.Information_item_id, rm_circulation.Info_item_subtype, rm_circulation.Physical_item_id, rm_circulation.Ppdm_guid, rm_circulation.Project_id, rm_circulation.Project_step_id, rm_circulation.Reference_num, rm_circulation.Remark, rm_circulation.Reserved_by, rm_circulation.Reserved_ind, rm_circulation.Source, rm_circulation.Status_date, rm_circulation.Row_changed_by, rm_circulation.Row_changed_date, rm_circulation.Row_created_by, rm_circulation.Row_effective_date, rm_circulation.Row_expiry_date, rm_circulation.Row_quality, rm_circulation.Circ_id)
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

func PatchRmCirculation(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_circulation set "
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
		} else if key == "circ_due_date" || key == "circ_in_date" || key == "circ_out_date" || key == "effective_date" || key == "expiry_date" || key == "status_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where circ_id = :id"

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

func DeleteRmCirculation(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_circulation dto.Rm_circulation
	rm_circulation.Circ_id = id

	stmt, err := db.Prepare("delete from rm_circulation where circ_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_circulation.Circ_id)
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


