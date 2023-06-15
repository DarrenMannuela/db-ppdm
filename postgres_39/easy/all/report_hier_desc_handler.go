package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetReportHierDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from report_hier_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Report_hier_desc

	for rows.Next() {
		var report_hier_desc dto.Report_hier_desc
		if err := rows.Scan(&report_hier_desc.Hierarchy_type_id, &report_hier_desc.Level_seq_no, &report_hier_desc.Active_ind, &report_hier_desc.Effective_date, &report_hier_desc.Expiry_date, &report_hier_desc.Level_name, &report_hier_desc.Level_type, &report_hier_desc.Ppdm_guid, &report_hier_desc.Remark, &report_hier_desc.Source, &report_hier_desc.System_id, &report_hier_desc.Table_name, &report_hier_desc.Row_changed_by, &report_hier_desc.Row_changed_date, &report_hier_desc.Row_created_by, &report_hier_desc.Row_created_date, &report_hier_desc.Row_effective_date, &report_hier_desc.Row_expiry_date, &report_hier_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append report_hier_desc to result
		result = append(result, report_hier_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetReportHierDesc(c *fiber.Ctx) error {
	var report_hier_desc dto.Report_hier_desc

	setDefaults(&report_hier_desc)

	if err := json.Unmarshal(c.Body(), &report_hier_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into report_hier_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	report_hier_desc.Row_created_date = formatDate(report_hier_desc.Row_created_date)
	report_hier_desc.Row_changed_date = formatDate(report_hier_desc.Row_changed_date)
	report_hier_desc.Effective_date = formatDateString(report_hier_desc.Effective_date)
	report_hier_desc.Expiry_date = formatDateString(report_hier_desc.Expiry_date)
	report_hier_desc.Row_effective_date = formatDateString(report_hier_desc.Row_effective_date)
	report_hier_desc.Row_expiry_date = formatDateString(report_hier_desc.Row_expiry_date)






	rows, err := stmt.Exec(report_hier_desc.Hierarchy_type_id, report_hier_desc.Level_seq_no, report_hier_desc.Active_ind, report_hier_desc.Effective_date, report_hier_desc.Expiry_date, report_hier_desc.Level_name, report_hier_desc.Level_type, report_hier_desc.Ppdm_guid, report_hier_desc.Remark, report_hier_desc.Source, report_hier_desc.System_id, report_hier_desc.Table_name, report_hier_desc.Row_changed_by, report_hier_desc.Row_changed_date, report_hier_desc.Row_created_by, report_hier_desc.Row_created_date, report_hier_desc.Row_effective_date, report_hier_desc.Row_expiry_date, report_hier_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateReportHierDesc(c *fiber.Ctx) error {
	var report_hier_desc dto.Report_hier_desc

	setDefaults(&report_hier_desc)

	if err := json.Unmarshal(c.Body(), &report_hier_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	report_hier_desc.Hierarchy_type_id = id

    if report_hier_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update report_hier_desc set level_seq_no = :1, active_ind = :2, effective_date = :3, expiry_date = :4, level_name = :5, level_type = :6, ppdm_guid = :7, remark = :8, source = :9, system_id = :10, table_name = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where hierarchy_type_id = :19")
	if err != nil {
		return err
	}

	report_hier_desc.Row_changed_date = formatDate(report_hier_desc.Row_changed_date)
	report_hier_desc.Effective_date = formatDateString(report_hier_desc.Effective_date)
	report_hier_desc.Expiry_date = formatDateString(report_hier_desc.Expiry_date)
	report_hier_desc.Row_effective_date = formatDateString(report_hier_desc.Row_effective_date)
	report_hier_desc.Row_expiry_date = formatDateString(report_hier_desc.Row_expiry_date)






	rows, err := stmt.Exec(report_hier_desc.Level_seq_no, report_hier_desc.Active_ind, report_hier_desc.Effective_date, report_hier_desc.Expiry_date, report_hier_desc.Level_name, report_hier_desc.Level_type, report_hier_desc.Ppdm_guid, report_hier_desc.Remark, report_hier_desc.Source, report_hier_desc.System_id, report_hier_desc.Table_name, report_hier_desc.Row_changed_by, report_hier_desc.Row_changed_date, report_hier_desc.Row_created_by, report_hier_desc.Row_effective_date, report_hier_desc.Row_expiry_date, report_hier_desc.Row_quality, report_hier_desc.Hierarchy_type_id)
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

func PatchReportHierDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update report_hier_desc set "
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

func DeleteReportHierDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var report_hier_desc dto.Report_hier_desc
	report_hier_desc.Hierarchy_type_id = id

	stmt, err := db.Prepare("delete from report_hier_desc where hierarchy_type_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(report_hier_desc.Hierarchy_type_id)
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


