package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetReportHierUse(c *fiber.Ctx) error {
	rows, err := db.Query("select * from report_hier_use")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Report_hier_use

	for rows.Next() {
		var report_hier_use dto.Report_hier_use
		if err := rows.Scan(&report_hier_use.Report_hierarchy_id, &report_hier_use.Component_id, &report_hier_use.Hierarchy_use_obs_no, &report_hier_use.Active_ind, &report_hier_use.Contribution_percent, &report_hier_use.Effective_date, &report_hier_use.Expiry_date, &report_hier_use.Pden_id, &report_hier_use.Pden_source, &report_hier_use.Pden_subtype, &report_hier_use.Ppdm_guid, &report_hier_use.Remark, &report_hier_use.Report_ind, &report_hier_use.Resent_id, &report_hier_use.Source, &report_hier_use.Substance_id, &report_hier_use.Row_changed_by, &report_hier_use.Row_changed_date, &report_hier_use.Row_created_by, &report_hier_use.Row_created_date, &report_hier_use.Row_effective_date, &report_hier_use.Row_expiry_date, &report_hier_use.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append report_hier_use to result
		result = append(result, report_hier_use)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetReportHierUse(c *fiber.Ctx) error {
	var report_hier_use dto.Report_hier_use

	setDefaults(&report_hier_use)

	if err := json.Unmarshal(c.Body(), &report_hier_use); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into report_hier_use values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	report_hier_use.Row_created_date = formatDate(report_hier_use.Row_created_date)
	report_hier_use.Row_changed_date = formatDate(report_hier_use.Row_changed_date)
	report_hier_use.Effective_date = formatDateString(report_hier_use.Effective_date)
	report_hier_use.Expiry_date = formatDateString(report_hier_use.Expiry_date)
	report_hier_use.Row_effective_date = formatDateString(report_hier_use.Row_effective_date)
	report_hier_use.Row_expiry_date = formatDateString(report_hier_use.Row_expiry_date)






	rows, err := stmt.Exec(report_hier_use.Report_hierarchy_id, report_hier_use.Component_id, report_hier_use.Hierarchy_use_obs_no, report_hier_use.Active_ind, report_hier_use.Contribution_percent, report_hier_use.Effective_date, report_hier_use.Expiry_date, report_hier_use.Pden_id, report_hier_use.Pden_source, report_hier_use.Pden_subtype, report_hier_use.Ppdm_guid, report_hier_use.Remark, report_hier_use.Report_ind, report_hier_use.Resent_id, report_hier_use.Source, report_hier_use.Substance_id, report_hier_use.Row_changed_by, report_hier_use.Row_changed_date, report_hier_use.Row_created_by, report_hier_use.Row_created_date, report_hier_use.Row_effective_date, report_hier_use.Row_expiry_date, report_hier_use.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateReportHierUse(c *fiber.Ctx) error {
	var report_hier_use dto.Report_hier_use

	setDefaults(&report_hier_use)

	if err := json.Unmarshal(c.Body(), &report_hier_use); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	report_hier_use.Report_hierarchy_id = id

    if report_hier_use.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update report_hier_use set component_id = :1, hierarchy_use_obs_no = :2, active_ind = :3, contribution_percent = :4, effective_date = :5, expiry_date = :6, pden_id = :7, pden_source = :8, pden_subtype = :9, ppdm_guid = :10, remark = :11, report_ind = :12, resent_id = :13, source = :14, substance_id = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where report_hierarchy_id = :23")
	if err != nil {
		return err
	}

	report_hier_use.Row_changed_date = formatDate(report_hier_use.Row_changed_date)
	report_hier_use.Effective_date = formatDateString(report_hier_use.Effective_date)
	report_hier_use.Expiry_date = formatDateString(report_hier_use.Expiry_date)
	report_hier_use.Row_effective_date = formatDateString(report_hier_use.Row_effective_date)
	report_hier_use.Row_expiry_date = formatDateString(report_hier_use.Row_expiry_date)






	rows, err := stmt.Exec(report_hier_use.Component_id, report_hier_use.Hierarchy_use_obs_no, report_hier_use.Active_ind, report_hier_use.Contribution_percent, report_hier_use.Effective_date, report_hier_use.Expiry_date, report_hier_use.Pden_id, report_hier_use.Pden_source, report_hier_use.Pden_subtype, report_hier_use.Ppdm_guid, report_hier_use.Remark, report_hier_use.Report_ind, report_hier_use.Resent_id, report_hier_use.Source, report_hier_use.Substance_id, report_hier_use.Row_changed_by, report_hier_use.Row_changed_date, report_hier_use.Row_created_by, report_hier_use.Row_effective_date, report_hier_use.Row_expiry_date, report_hier_use.Row_quality, report_hier_use.Report_hierarchy_id)
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

func PatchReportHierUse(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update report_hier_use set "
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

func DeleteReportHierUse(c *fiber.Ctx) error {
	id := c.Params("id")
	var report_hier_use dto.Report_hier_use
	report_hier_use.Report_hierarchy_id = id

	stmt, err := db.Prepare("delete from report_hier_use where report_hierarchy_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(report_hier_use.Report_hierarchy_id)
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


