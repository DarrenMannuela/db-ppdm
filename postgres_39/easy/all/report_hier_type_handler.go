package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetReportHierType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from report_hier_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Report_hier_type

	for rows.Next() {
		var report_hier_type dto.Report_hier_type
		if err := rows.Scan(&report_hier_type.Hierarchy_type_id, &report_hier_type.Active_ind, &report_hier_type.Effective_date, &report_hier_type.Expiry_date, &report_hier_type.Hierarchy_name, &report_hier_type.Hierarchy_type, &report_hier_type.Ppdm_guid, &report_hier_type.Remark, &report_hier_type.Source, &report_hier_type.Row_changed_by, &report_hier_type.Row_changed_date, &report_hier_type.Row_created_by, &report_hier_type.Row_created_date, &report_hier_type.Row_effective_date, &report_hier_type.Row_expiry_date, &report_hier_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append report_hier_type to result
		result = append(result, report_hier_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetReportHierType(c *fiber.Ctx) error {
	var report_hier_type dto.Report_hier_type

	setDefaults(&report_hier_type)

	if err := json.Unmarshal(c.Body(), &report_hier_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into report_hier_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	report_hier_type.Row_created_date = formatDate(report_hier_type.Row_created_date)
	report_hier_type.Row_changed_date = formatDate(report_hier_type.Row_changed_date)
	report_hier_type.Effective_date = formatDateString(report_hier_type.Effective_date)
	report_hier_type.Expiry_date = formatDateString(report_hier_type.Expiry_date)
	report_hier_type.Row_effective_date = formatDateString(report_hier_type.Row_effective_date)
	report_hier_type.Row_expiry_date = formatDateString(report_hier_type.Row_expiry_date)






	rows, err := stmt.Exec(report_hier_type.Hierarchy_type_id, report_hier_type.Active_ind, report_hier_type.Effective_date, report_hier_type.Expiry_date, report_hier_type.Hierarchy_name, report_hier_type.Hierarchy_type, report_hier_type.Ppdm_guid, report_hier_type.Remark, report_hier_type.Source, report_hier_type.Row_changed_by, report_hier_type.Row_changed_date, report_hier_type.Row_created_by, report_hier_type.Row_created_date, report_hier_type.Row_effective_date, report_hier_type.Row_expiry_date, report_hier_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateReportHierType(c *fiber.Ctx) error {
	var report_hier_type dto.Report_hier_type

	setDefaults(&report_hier_type)

	if err := json.Unmarshal(c.Body(), &report_hier_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	report_hier_type.Hierarchy_type_id = id

    if report_hier_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update report_hier_type set active_ind = :1, effective_date = :2, expiry_date = :3, hierarchy_name = :4, hierarchy_type = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where hierarchy_type_id = :16")
	if err != nil {
		return err
	}

	report_hier_type.Row_changed_date = formatDate(report_hier_type.Row_changed_date)
	report_hier_type.Effective_date = formatDateString(report_hier_type.Effective_date)
	report_hier_type.Expiry_date = formatDateString(report_hier_type.Expiry_date)
	report_hier_type.Row_effective_date = formatDateString(report_hier_type.Row_effective_date)
	report_hier_type.Row_expiry_date = formatDateString(report_hier_type.Row_expiry_date)






	rows, err := stmt.Exec(report_hier_type.Active_ind, report_hier_type.Effective_date, report_hier_type.Expiry_date, report_hier_type.Hierarchy_name, report_hier_type.Hierarchy_type, report_hier_type.Ppdm_guid, report_hier_type.Remark, report_hier_type.Source, report_hier_type.Row_changed_by, report_hier_type.Row_changed_date, report_hier_type.Row_created_by, report_hier_type.Row_effective_date, report_hier_type.Row_expiry_date, report_hier_type.Row_quality, report_hier_type.Hierarchy_type_id)
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

func PatchReportHierType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update report_hier_type set "
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
	query += " where hierarchy_type_id = :id"

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

func DeleteReportHierType(c *fiber.Ctx) error {
	id := c.Params("id")
	var report_hier_type dto.Report_hier_type
	report_hier_type.Hierarchy_type_id = id

	stmt, err := db.Prepare("delete from report_hier_type where hierarchy_type_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(report_hier_type.Hierarchy_type_id)
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


