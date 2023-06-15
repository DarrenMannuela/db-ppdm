package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWorkOrderAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from work_order_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Work_order_alias

	for rows.Next() {
		var work_order_alias dto.Work_order_alias
		if err := rows.Scan(&work_order_alias.Work_order_id, &work_order_alias.Alias_id, &work_order_alias.Abbreviation, &work_order_alias.Active_ind, &work_order_alias.Alias_long_name, &work_order_alias.Alias_reason_type, &work_order_alias.Alias_short_name, &work_order_alias.Alias_type, &work_order_alias.Amended_date, &work_order_alias.Created_date, &work_order_alias.Effective_date, &work_order_alias.Expiry_date, &work_order_alias.Work_order_id, &work_order_alias.Original_ind, &work_order_alias.Owner_ba_id, &work_order_alias.Ppdm_guid, &work_order_alias.Preferred_ind, &work_order_alias.Reason_desc, &work_order_alias.Remark, &work_order_alias.Source, &work_order_alias.Source_document_id, &work_order_alias.Struckoff_date, &work_order_alias.Sw_application_id, &work_order_alias.Use_rule, &work_order_alias.Row_changed_by, &work_order_alias.Row_changed_date, &work_order_alias.Row_created_by, &work_order_alias.Row_created_date, &work_order_alias.Row_effective_date, &work_order_alias.Row_expiry_date, &work_order_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append work_order_alias to result
		result = append(result, work_order_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWorkOrderAlias(c *fiber.Ctx) error {
	var work_order_alias dto.Work_order_alias

	setDefaults(&work_order_alias)

	if err := json.Unmarshal(c.Body(), &work_order_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into work_order_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	work_order_alias.Row_created_date = formatDate(work_order_alias.Row_created_date)
	work_order_alias.Row_changed_date = formatDate(work_order_alias.Row_changed_date)
	work_order_alias.Amended_date = formatDateString(work_order_alias.Amended_date)
	work_order_alias.Created_date = formatDateString(work_order_alias.Created_date)
	work_order_alias.Effective_date = formatDateString(work_order_alias.Effective_date)
	work_order_alias.Expiry_date = formatDateString(work_order_alias.Expiry_date)
	work_order_alias.Struckoff_date = formatDateString(work_order_alias.Struckoff_date)
	work_order_alias.Row_effective_date = formatDateString(work_order_alias.Row_effective_date)
	work_order_alias.Row_expiry_date = formatDateString(work_order_alias.Row_expiry_date)






	rows, err := stmt.Exec(work_order_alias.Work_order_id, work_order_alias.Alias_id, work_order_alias.Abbreviation, work_order_alias.Active_ind, work_order_alias.Alias_long_name, work_order_alias.Alias_reason_type, work_order_alias.Alias_short_name, work_order_alias.Alias_type, work_order_alias.Amended_date, work_order_alias.Created_date, work_order_alias.Effective_date, work_order_alias.Expiry_date, work_order_alias.Work_order_id, work_order_alias.Original_ind, work_order_alias.Owner_ba_id, work_order_alias.Ppdm_guid, work_order_alias.Preferred_ind, work_order_alias.Reason_desc, work_order_alias.Remark, work_order_alias.Source, work_order_alias.Source_document_id, work_order_alias.Struckoff_date, work_order_alias.Sw_application_id, work_order_alias.Use_rule, work_order_alias.Row_changed_by, work_order_alias.Row_changed_date, work_order_alias.Row_created_by, work_order_alias.Row_created_date, work_order_alias.Row_effective_date, work_order_alias.Row_expiry_date, work_order_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWorkOrderAlias(c *fiber.Ctx) error {
	var work_order_alias dto.Work_order_alias

	setDefaults(&work_order_alias)

	if err := json.Unmarshal(c.Body(), &work_order_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	work_order_alias.Work_order_id = id

    if work_order_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update work_order_alias set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, work_order_id = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where work_order_id = :31")
	if err != nil {
		return err
	}

	work_order_alias.Row_changed_date = formatDate(work_order_alias.Row_changed_date)
	work_order_alias.Amended_date = formatDateString(work_order_alias.Amended_date)
	work_order_alias.Created_date = formatDateString(work_order_alias.Created_date)
	work_order_alias.Effective_date = formatDateString(work_order_alias.Effective_date)
	work_order_alias.Expiry_date = formatDateString(work_order_alias.Expiry_date)
	work_order_alias.Struckoff_date = formatDateString(work_order_alias.Struckoff_date)
	work_order_alias.Row_effective_date = formatDateString(work_order_alias.Row_effective_date)
	work_order_alias.Row_expiry_date = formatDateString(work_order_alias.Row_expiry_date)






	rows, err := stmt.Exec(work_order_alias.Alias_id, work_order_alias.Abbreviation, work_order_alias.Active_ind, work_order_alias.Alias_long_name, work_order_alias.Alias_reason_type, work_order_alias.Alias_short_name, work_order_alias.Alias_type, work_order_alias.Amended_date, work_order_alias.Created_date, work_order_alias.Effective_date, work_order_alias.Expiry_date, work_order_alias.Work_order_id, work_order_alias.Original_ind, work_order_alias.Owner_ba_id, work_order_alias.Ppdm_guid, work_order_alias.Preferred_ind, work_order_alias.Reason_desc, work_order_alias.Remark, work_order_alias.Source, work_order_alias.Source_document_id, work_order_alias.Struckoff_date, work_order_alias.Sw_application_id, work_order_alias.Use_rule, work_order_alias.Row_changed_by, work_order_alias.Row_changed_date, work_order_alias.Row_created_by, work_order_alias.Row_effective_date, work_order_alias.Row_expiry_date, work_order_alias.Row_quality, work_order_alias.Work_order_id)
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

func PatchWorkOrderAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update work_order_alias set "
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
		} else if key == "amended_date" || key == "created_date" || key == "effective_date" || key == "expiry_date" || key == "struckoff_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWorkOrderAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var work_order_alias dto.Work_order_alias
	work_order_alias.Work_order_id = id

	stmt, err := db.Prepare("delete from work_order_alias where work_order_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(work_order_alias.Work_order_id)
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


