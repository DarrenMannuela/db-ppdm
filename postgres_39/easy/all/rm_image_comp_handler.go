package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmImageComp(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_image_comp")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_image_comp

	for rows.Next() {
		var rm_image_comp dto.Rm_image_comp
		if err := rows.Scan(&rm_image_comp.Physical_item_id, &rm_image_comp.Log_section_id_1, &rm_image_comp.Log_section_id_2, &rm_image_comp.Active_ind, &rm_image_comp.Composite_seq_no, &rm_image_comp.Effective_date, &rm_image_comp.Expiry_date, &rm_image_comp.Ppdm_guid, &rm_image_comp.Remark, &rm_image_comp.Source, &rm_image_comp.Row_changed_by, &rm_image_comp.Row_changed_date, &rm_image_comp.Row_created_by, &rm_image_comp.Row_created_date, &rm_image_comp.Row_effective_date, &rm_image_comp.Row_expiry_date, &rm_image_comp.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_image_comp to result
		result = append(result, rm_image_comp)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmImageComp(c *fiber.Ctx) error {
	var rm_image_comp dto.Rm_image_comp

	setDefaults(&rm_image_comp)

	if err := json.Unmarshal(c.Body(), &rm_image_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_image_comp values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	rm_image_comp.Row_created_date = formatDate(rm_image_comp.Row_created_date)
	rm_image_comp.Row_changed_date = formatDate(rm_image_comp.Row_changed_date)
	rm_image_comp.Effective_date = formatDateString(rm_image_comp.Effective_date)
	rm_image_comp.Expiry_date = formatDateString(rm_image_comp.Expiry_date)
	rm_image_comp.Row_effective_date = formatDateString(rm_image_comp.Row_effective_date)
	rm_image_comp.Row_expiry_date = formatDateString(rm_image_comp.Row_expiry_date)






	rows, err := stmt.Exec(rm_image_comp.Physical_item_id, rm_image_comp.Log_section_id_1, rm_image_comp.Log_section_id_2, rm_image_comp.Active_ind, rm_image_comp.Composite_seq_no, rm_image_comp.Effective_date, rm_image_comp.Expiry_date, rm_image_comp.Ppdm_guid, rm_image_comp.Remark, rm_image_comp.Source, rm_image_comp.Row_changed_by, rm_image_comp.Row_changed_date, rm_image_comp.Row_created_by, rm_image_comp.Row_created_date, rm_image_comp.Row_effective_date, rm_image_comp.Row_expiry_date, rm_image_comp.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmImageComp(c *fiber.Ctx) error {
	var rm_image_comp dto.Rm_image_comp

	setDefaults(&rm_image_comp)

	if err := json.Unmarshal(c.Body(), &rm_image_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_image_comp.Physical_item_id = id

    if rm_image_comp.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_image_comp set log_section_id_1 = :1, log_section_id_2 = :2, active_ind = :3, composite_seq_no = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where physical_item_id = :17")
	if err != nil {
		return err
	}

	rm_image_comp.Row_changed_date = formatDate(rm_image_comp.Row_changed_date)
	rm_image_comp.Effective_date = formatDateString(rm_image_comp.Effective_date)
	rm_image_comp.Expiry_date = formatDateString(rm_image_comp.Expiry_date)
	rm_image_comp.Row_effective_date = formatDateString(rm_image_comp.Row_effective_date)
	rm_image_comp.Row_expiry_date = formatDateString(rm_image_comp.Row_expiry_date)






	rows, err := stmt.Exec(rm_image_comp.Log_section_id_1, rm_image_comp.Log_section_id_2, rm_image_comp.Active_ind, rm_image_comp.Composite_seq_no, rm_image_comp.Effective_date, rm_image_comp.Expiry_date, rm_image_comp.Ppdm_guid, rm_image_comp.Remark, rm_image_comp.Source, rm_image_comp.Row_changed_by, rm_image_comp.Row_changed_date, rm_image_comp.Row_created_by, rm_image_comp.Row_effective_date, rm_image_comp.Row_expiry_date, rm_image_comp.Row_quality, rm_image_comp.Physical_item_id)
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

func PatchRmImageComp(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_image_comp set "
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
	query += " where physical_item_id = :id"

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

func DeleteRmImageComp(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_image_comp dto.Rm_image_comp
	rm_image_comp.Physical_item_id = id

	stmt, err := db.Prepare("delete from rm_image_comp where physical_item_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_image_comp.Physical_item_id)
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


