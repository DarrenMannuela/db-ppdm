package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetReportHierLevel(c *fiber.Ctx) error {
	rows, err := db.Query("select * from report_hier_level")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Report_hier_level

	for rows.Next() {
		var report_hier_level dto.Report_hier_level
		if err := rows.Scan(&report_hier_level.Report_hierarchy_id, &report_hier_level.Component_id, &report_hier_level.Active_ind, &report_hier_level.Area_id, &report_hier_level.Area_type, &report_hier_level.Ba_organization_id, &report_hier_level.Ba_organization_seq_no, &report_hier_level.Business_associate_id, &report_hier_level.Component_type, &report_hier_level.Contract_id, &report_hier_level.Effective_date, &report_hier_level.Equipment_id, &report_hier_level.Expiry_date, &report_hier_level.Facility_id, &report_hier_level.Facility_type, &report_hier_level.Field_id, &report_hier_level.Finance_id, &report_hier_level.Hierarchy_type_id, &report_hier_level.Land_right_id, &report_hier_level.Land_right_subtype, &report_hier_level.Level_seq_no, &report_hier_level.Parent_component_id, &report_hier_level.Pool_id, &report_hier_level.Ppdm_guid, &report_hier_level.Project_id, &report_hier_level.Remark, &report_hier_level.Sf_subtype, &report_hier_level.Source, &report_hier_level.Strat_name_set_id, &report_hier_level.Strat_unit_id, &report_hier_level.Support_facility_id, &report_hier_level.Uwi, &report_hier_level.Row_changed_by, &report_hier_level.Row_changed_date, &report_hier_level.Row_created_by, &report_hier_level.Row_created_date, &report_hier_level.Row_effective_date, &report_hier_level.Row_expiry_date, &report_hier_level.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append report_hier_level to result
		result = append(result, report_hier_level)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetReportHierLevel(c *fiber.Ctx) error {
	var report_hier_level dto.Report_hier_level

	setDefaults(&report_hier_level)

	if err := json.Unmarshal(c.Body(), &report_hier_level); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into report_hier_level values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	report_hier_level.Row_created_date = formatDate(report_hier_level.Row_created_date)
	report_hier_level.Row_changed_date = formatDate(report_hier_level.Row_changed_date)
	report_hier_level.Effective_date = formatDateString(report_hier_level.Effective_date)
	report_hier_level.Expiry_date = formatDateString(report_hier_level.Expiry_date)
	report_hier_level.Row_effective_date = formatDateString(report_hier_level.Row_effective_date)
	report_hier_level.Row_expiry_date = formatDateString(report_hier_level.Row_expiry_date)






	rows, err := stmt.Exec(report_hier_level.Report_hierarchy_id, report_hier_level.Component_id, report_hier_level.Active_ind, report_hier_level.Area_id, report_hier_level.Area_type, report_hier_level.Ba_organization_id, report_hier_level.Ba_organization_seq_no, report_hier_level.Business_associate_id, report_hier_level.Component_type, report_hier_level.Contract_id, report_hier_level.Effective_date, report_hier_level.Equipment_id, report_hier_level.Expiry_date, report_hier_level.Facility_id, report_hier_level.Facility_type, report_hier_level.Field_id, report_hier_level.Finance_id, report_hier_level.Hierarchy_type_id, report_hier_level.Land_right_id, report_hier_level.Land_right_subtype, report_hier_level.Level_seq_no, report_hier_level.Parent_component_id, report_hier_level.Pool_id, report_hier_level.Ppdm_guid, report_hier_level.Project_id, report_hier_level.Remark, report_hier_level.Sf_subtype, report_hier_level.Source, report_hier_level.Strat_name_set_id, report_hier_level.Strat_unit_id, report_hier_level.Support_facility_id, report_hier_level.Uwi, report_hier_level.Row_changed_by, report_hier_level.Row_changed_date, report_hier_level.Row_created_by, report_hier_level.Row_created_date, report_hier_level.Row_effective_date, report_hier_level.Row_expiry_date, report_hier_level.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateReportHierLevel(c *fiber.Ctx) error {
	var report_hier_level dto.Report_hier_level

	setDefaults(&report_hier_level)

	if err := json.Unmarshal(c.Body(), &report_hier_level); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	report_hier_level.Report_hierarchy_id = id

    if report_hier_level.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update report_hier_level set component_id = :1, active_ind = :2, area_id = :3, area_type = :4, ba_organization_id = :5, ba_organization_seq_no = :6, business_associate_id = :7, component_type = :8, contract_id = :9, effective_date = :10, equipment_id = :11, expiry_date = :12, facility_id = :13, facility_type = :14, field_id = :15, finance_id = :16, hierarchy_type_id = :17, land_right_id = :18, land_right_subtype = :19, level_seq_no = :20, parent_component_id = :21, pool_id = :22, ppdm_guid = :23, project_id = :24, remark = :25, sf_subtype = :26, source = :27, strat_name_set_id = :28, strat_unit_id = :29, support_facility_id = :30, uwi = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where report_hierarchy_id = :39")
	if err != nil {
		return err
	}

	report_hier_level.Row_changed_date = formatDate(report_hier_level.Row_changed_date)
	report_hier_level.Effective_date = formatDateString(report_hier_level.Effective_date)
	report_hier_level.Expiry_date = formatDateString(report_hier_level.Expiry_date)
	report_hier_level.Row_effective_date = formatDateString(report_hier_level.Row_effective_date)
	report_hier_level.Row_expiry_date = formatDateString(report_hier_level.Row_expiry_date)






	rows, err := stmt.Exec(report_hier_level.Component_id, report_hier_level.Active_ind, report_hier_level.Area_id, report_hier_level.Area_type, report_hier_level.Ba_organization_id, report_hier_level.Ba_organization_seq_no, report_hier_level.Business_associate_id, report_hier_level.Component_type, report_hier_level.Contract_id, report_hier_level.Effective_date, report_hier_level.Equipment_id, report_hier_level.Expiry_date, report_hier_level.Facility_id, report_hier_level.Facility_type, report_hier_level.Field_id, report_hier_level.Finance_id, report_hier_level.Hierarchy_type_id, report_hier_level.Land_right_id, report_hier_level.Land_right_subtype, report_hier_level.Level_seq_no, report_hier_level.Parent_component_id, report_hier_level.Pool_id, report_hier_level.Ppdm_guid, report_hier_level.Project_id, report_hier_level.Remark, report_hier_level.Sf_subtype, report_hier_level.Source, report_hier_level.Strat_name_set_id, report_hier_level.Strat_unit_id, report_hier_level.Support_facility_id, report_hier_level.Uwi, report_hier_level.Row_changed_by, report_hier_level.Row_changed_date, report_hier_level.Row_created_by, report_hier_level.Row_effective_date, report_hier_level.Row_expiry_date, report_hier_level.Row_quality, report_hier_level.Report_hierarchy_id)
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

func PatchReportHierLevel(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update report_hier_level set "
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

func DeleteReportHierLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	var report_hier_level dto.Report_hier_level
	report_hier_level.Report_hierarchy_id = id

	stmt, err := db.Prepare("delete from report_hier_level where report_hierarchy_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(report_hier_level.Report_hierarchy_id)
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


