package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetReportHier(c *fiber.Ctx) error {
	rows, err := db.Query("select * from report_hier")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Report_hier

	for rows.Next() {
		var report_hier dto.Report_hier
		if err := rows.Scan(&report_hier.Report_hierarchy_id, &report_hier.Active_ind, &report_hier.Effective_date, &report_hier.Expiry_date, &report_hier.Hierarchy_type, &report_hier.Ppdm_guid, &report_hier.Preferred_name, &report_hier.Remark, &report_hier.Source, &report_hier.Row_changed_by, &report_hier.Row_changed_date, &report_hier.Row_created_by, &report_hier.Row_created_date, &report_hier.Row_effective_date, &report_hier.Row_expiry_date, &report_hier.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append report_hier to result
		result = append(result, report_hier)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetReportHier(c *fiber.Ctx) error {
	var report_hier dto.Report_hier

	setDefaults(&report_hier)

	if err := json.Unmarshal(c.Body(), &report_hier); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into report_hier values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	report_hier.Row_created_date = formatDate(report_hier.Row_created_date)
	report_hier.Row_changed_date = formatDate(report_hier.Row_changed_date)
	report_hier.Effective_date = formatDateString(report_hier.Effective_date)
	report_hier.Expiry_date = formatDateString(report_hier.Expiry_date)
	report_hier.Row_effective_date = formatDateString(report_hier.Row_effective_date)
	report_hier.Row_expiry_date = formatDateString(report_hier.Row_expiry_date)






	rows, err := stmt.Exec(report_hier.Report_hierarchy_id, report_hier.Active_ind, report_hier.Effective_date, report_hier.Expiry_date, report_hier.Hierarchy_type, report_hier.Ppdm_guid, report_hier.Preferred_name, report_hier.Remark, report_hier.Source, report_hier.Row_changed_by, report_hier.Row_changed_date, report_hier.Row_created_by, report_hier.Row_created_date, report_hier.Row_effective_date, report_hier.Row_expiry_date, report_hier.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateReportHier(c *fiber.Ctx) error {
	var report_hier dto.Report_hier

	setDefaults(&report_hier)

	if err := json.Unmarshal(c.Body(), &report_hier); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	report_hier.Report_hierarchy_id = id

    if report_hier.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update report_hier set active_ind = :1, effective_date = :2, expiry_date = :3, hierarchy_type = :4, ppdm_guid = :5, preferred_name = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where report_hierarchy_id = :16")
	if err != nil {
		return err
	}

	report_hier.Row_changed_date = formatDate(report_hier.Row_changed_date)
	report_hier.Effective_date = formatDateString(report_hier.Effective_date)
	report_hier.Expiry_date = formatDateString(report_hier.Expiry_date)
	report_hier.Row_effective_date = formatDateString(report_hier.Row_effective_date)
	report_hier.Row_expiry_date = formatDateString(report_hier.Row_expiry_date)






	rows, err := stmt.Exec(report_hier.Active_ind, report_hier.Effective_date, report_hier.Expiry_date, report_hier.Hierarchy_type, report_hier.Ppdm_guid, report_hier.Preferred_name, report_hier.Remark, report_hier.Source, report_hier.Row_changed_by, report_hier.Row_changed_date, report_hier.Row_created_by, report_hier.Row_effective_date, report_hier.Row_expiry_date, report_hier.Row_quality, report_hier.Report_hierarchy_id)
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

func PatchReportHier(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update report_hier set "
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
	query += " where report_hierarchy_id = :id"

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

func DeleteReportHier(c *fiber.Ctx) error {
	id := c.Params("id")
	var report_hier dto.Report_hier
	report_hier.Report_hierarchy_id = id

	stmt, err := db.Prepare("delete from report_hier where report_hierarchy_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(report_hier.Report_hierarchy_id)
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


