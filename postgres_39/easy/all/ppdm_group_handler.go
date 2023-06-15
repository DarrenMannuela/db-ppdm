package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmGroup(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_group")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_group

	for rows.Next() {
		var ppdm_group dto.Ppdm_group
		if err := rows.Scan(&ppdm_group.Group_id, &ppdm_group.Active_ind, &ppdm_group.Default_unit_system_id, &ppdm_group.Effective_date, &ppdm_group.Expiry_date, &ppdm_group.Group_name, &ppdm_group.Group_type, &ppdm_group.Ppdm_guid, &ppdm_group.Remark, &ppdm_group.Report_heading1, &ppdm_group.Report_heading2, &ppdm_group.Screen_heading1, &ppdm_group.Screen_heading2, &ppdm_group.Source, &ppdm_group.Sw_application_id, &ppdm_group.Row_changed_by, &ppdm_group.Row_changed_date, &ppdm_group.Row_created_by, &ppdm_group.Row_created_date, &ppdm_group.Row_effective_date, &ppdm_group.Row_expiry_date, &ppdm_group.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_group to result
		result = append(result, ppdm_group)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmGroup(c *fiber.Ctx) error {
	var ppdm_group dto.Ppdm_group

	setDefaults(&ppdm_group)

	if err := json.Unmarshal(c.Body(), &ppdm_group); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_group values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	ppdm_group.Row_created_date = formatDate(ppdm_group.Row_created_date)
	ppdm_group.Row_changed_date = formatDate(ppdm_group.Row_changed_date)
	ppdm_group.Effective_date = formatDateString(ppdm_group.Effective_date)
	ppdm_group.Expiry_date = formatDateString(ppdm_group.Expiry_date)
	ppdm_group.Row_effective_date = formatDateString(ppdm_group.Row_effective_date)
	ppdm_group.Row_expiry_date = formatDateString(ppdm_group.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_group.Group_id, ppdm_group.Active_ind, ppdm_group.Default_unit_system_id, ppdm_group.Effective_date, ppdm_group.Expiry_date, ppdm_group.Group_name, ppdm_group.Group_type, ppdm_group.Ppdm_guid, ppdm_group.Remark, ppdm_group.Report_heading1, ppdm_group.Report_heading2, ppdm_group.Screen_heading1, ppdm_group.Screen_heading2, ppdm_group.Source, ppdm_group.Sw_application_id, ppdm_group.Row_changed_by, ppdm_group.Row_changed_date, ppdm_group.Row_created_by, ppdm_group.Row_created_date, ppdm_group.Row_effective_date, ppdm_group.Row_expiry_date, ppdm_group.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmGroup(c *fiber.Ctx) error {
	var ppdm_group dto.Ppdm_group

	setDefaults(&ppdm_group)

	if err := json.Unmarshal(c.Body(), &ppdm_group); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_group.Group_id = id

    if ppdm_group.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_group set active_ind = :1, default_unit_system_id = :2, effective_date = :3, expiry_date = :4, group_name = :5, group_type = :6, ppdm_guid = :7, remark = :8, report_heading1 = :9, report_heading2 = :10, screen_heading1 = :11, screen_heading2 = :12, source = :13, sw_application_id = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where group_id = :22")
	if err != nil {
		return err
	}

	ppdm_group.Row_changed_date = formatDate(ppdm_group.Row_changed_date)
	ppdm_group.Effective_date = formatDateString(ppdm_group.Effective_date)
	ppdm_group.Expiry_date = formatDateString(ppdm_group.Expiry_date)
	ppdm_group.Row_effective_date = formatDateString(ppdm_group.Row_effective_date)
	ppdm_group.Row_expiry_date = formatDateString(ppdm_group.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_group.Active_ind, ppdm_group.Default_unit_system_id, ppdm_group.Effective_date, ppdm_group.Expiry_date, ppdm_group.Group_name, ppdm_group.Group_type, ppdm_group.Ppdm_guid, ppdm_group.Remark, ppdm_group.Report_heading1, ppdm_group.Report_heading2, ppdm_group.Screen_heading1, ppdm_group.Screen_heading2, ppdm_group.Source, ppdm_group.Sw_application_id, ppdm_group.Row_changed_by, ppdm_group.Row_changed_date, ppdm_group.Row_created_by, ppdm_group.Row_effective_date, ppdm_group.Row_expiry_date, ppdm_group.Row_quality, ppdm_group.Group_id)
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

func PatchPpdmGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_group set "
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
	query += " where group_id = :id"

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

func DeletePpdmGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_group dto.Ppdm_group
	ppdm_group.Group_id = id

	stmt, err := db.Prepare("delete from ppdm_group where group_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_group.Group_id)
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


