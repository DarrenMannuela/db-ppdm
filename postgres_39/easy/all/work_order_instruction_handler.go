package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWorkOrderInstruction(c *fiber.Ctx) error {
	rows, err := db.Query("select * from work_order_instruction")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Work_order_instruction

	for rows.Next() {
		var work_order_instruction dto.Work_order_instruction
		if err := rows.Scan(&work_order_instruction.Work_order_id, &work_order_instruction.Instruction_id, &work_order_instruction.Active_ind, &work_order_instruction.Ba_address_obs_no, &work_order_instruction.Ba_address_source, &work_order_instruction.Ba_obs_no, &work_order_instruction.Business_associate_id, &work_order_instruction.Effective_date, &work_order_instruction.Expiry_date, &work_order_instruction.Instruction_desc, &work_order_instruction.Instruction_type, &work_order_instruction.Instruction_value, &work_order_instruction.Instruction_value_ouom, &work_order_instruction.Instruction_value_uom, &work_order_instruction.Ppdm_guid, &work_order_instruction.Remark, &work_order_instruction.Source, &work_order_instruction.Row_changed_by, &work_order_instruction.Row_changed_date, &work_order_instruction.Row_created_by, &work_order_instruction.Row_created_date, &work_order_instruction.Row_effective_date, &work_order_instruction.Row_expiry_date, &work_order_instruction.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append work_order_instruction to result
		result = append(result, work_order_instruction)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWorkOrderInstruction(c *fiber.Ctx) error {
	var work_order_instruction dto.Work_order_instruction

	setDefaults(&work_order_instruction)

	if err := json.Unmarshal(c.Body(), &work_order_instruction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into work_order_instruction values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	work_order_instruction.Row_created_date = formatDate(work_order_instruction.Row_created_date)
	work_order_instruction.Row_changed_date = formatDate(work_order_instruction.Row_changed_date)
	work_order_instruction.Effective_date = formatDateString(work_order_instruction.Effective_date)
	work_order_instruction.Expiry_date = formatDateString(work_order_instruction.Expiry_date)
	work_order_instruction.Row_effective_date = formatDateString(work_order_instruction.Row_effective_date)
	work_order_instruction.Row_expiry_date = formatDateString(work_order_instruction.Row_expiry_date)






	rows, err := stmt.Exec(work_order_instruction.Work_order_id, work_order_instruction.Instruction_id, work_order_instruction.Active_ind, work_order_instruction.Ba_address_obs_no, work_order_instruction.Ba_address_source, work_order_instruction.Ba_obs_no, work_order_instruction.Business_associate_id, work_order_instruction.Effective_date, work_order_instruction.Expiry_date, work_order_instruction.Instruction_desc, work_order_instruction.Instruction_type, work_order_instruction.Instruction_value, work_order_instruction.Instruction_value_ouom, work_order_instruction.Instruction_value_uom, work_order_instruction.Ppdm_guid, work_order_instruction.Remark, work_order_instruction.Source, work_order_instruction.Row_changed_by, work_order_instruction.Row_changed_date, work_order_instruction.Row_created_by, work_order_instruction.Row_created_date, work_order_instruction.Row_effective_date, work_order_instruction.Row_expiry_date, work_order_instruction.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWorkOrderInstruction(c *fiber.Ctx) error {
	var work_order_instruction dto.Work_order_instruction

	setDefaults(&work_order_instruction)

	if err := json.Unmarshal(c.Body(), &work_order_instruction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	work_order_instruction.Work_order_id = id

    if work_order_instruction.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update work_order_instruction set instruction_id = :1, active_ind = :2, ba_address_obs_no = :3, ba_address_source = :4, ba_obs_no = :5, business_associate_id = :6, effective_date = :7, expiry_date = :8, instruction_desc = :9, instruction_type = :10, instruction_value = :11, instruction_value_ouom = :12, instruction_value_uom = :13, ppdm_guid = :14, remark = :15, source = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where work_order_id = :24")
	if err != nil {
		return err
	}

	work_order_instruction.Row_changed_date = formatDate(work_order_instruction.Row_changed_date)
	work_order_instruction.Effective_date = formatDateString(work_order_instruction.Effective_date)
	work_order_instruction.Expiry_date = formatDateString(work_order_instruction.Expiry_date)
	work_order_instruction.Row_effective_date = formatDateString(work_order_instruction.Row_effective_date)
	work_order_instruction.Row_expiry_date = formatDateString(work_order_instruction.Row_expiry_date)






	rows, err := stmt.Exec(work_order_instruction.Instruction_id, work_order_instruction.Active_ind, work_order_instruction.Ba_address_obs_no, work_order_instruction.Ba_address_source, work_order_instruction.Ba_obs_no, work_order_instruction.Business_associate_id, work_order_instruction.Effective_date, work_order_instruction.Expiry_date, work_order_instruction.Instruction_desc, work_order_instruction.Instruction_type, work_order_instruction.Instruction_value, work_order_instruction.Instruction_value_ouom, work_order_instruction.Instruction_value_uom, work_order_instruction.Ppdm_guid, work_order_instruction.Remark, work_order_instruction.Source, work_order_instruction.Row_changed_by, work_order_instruction.Row_changed_date, work_order_instruction.Row_created_by, work_order_instruction.Row_effective_date, work_order_instruction.Row_expiry_date, work_order_instruction.Row_quality, work_order_instruction.Work_order_id)
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

func PatchWorkOrderInstruction(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update work_order_instruction set "
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
	query += " where work_order_id = :id"

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

func DeleteWorkOrderInstruction(c *fiber.Ctx) error {
	id := c.Params("id")
	var work_order_instruction dto.Work_order_instruction
	work_order_instruction.Work_order_id = id

	stmt, err := db.Prepare("delete from work_order_instruction where work_order_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(work_order_instruction.Work_order_id)
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


