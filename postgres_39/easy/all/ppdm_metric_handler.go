package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmMetric(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_metric")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_metric

	for rows.Next() {
		var ppdm_metric dto.Ppdm_metric
		if err := rows.Scan(&ppdm_metric.Metric_id, &ppdm_metric.Active_ind, &ppdm_metric.Effective_date, &ppdm_metric.End_date, &ppdm_metric.Expiry_date, &ppdm_metric.Metric_name, &ppdm_metric.Metric_procedure, &ppdm_metric.Metric_type, &ppdm_metric.Organization_id, &ppdm_metric.Organization_seq_no, &ppdm_metric.Owner_ba_id, &ppdm_metric.Ppdm_guid, &ppdm_metric.Procedure_id, &ppdm_metric.Procedure_system_id, &ppdm_metric.Projected_final_count, &ppdm_metric.Purpose_desc, &ppdm_metric.Remark, &ppdm_metric.Source, &ppdm_metric.Start_date, &ppdm_metric.Row_changed_by, &ppdm_metric.Row_changed_date, &ppdm_metric.Row_created_by, &ppdm_metric.Row_created_date, &ppdm_metric.Row_effective_date, &ppdm_metric.Row_expiry_date, &ppdm_metric.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_metric to result
		result = append(result, ppdm_metric)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmMetric(c *fiber.Ctx) error {
	var ppdm_metric dto.Ppdm_metric

	setDefaults(&ppdm_metric)

	if err := json.Unmarshal(c.Body(), &ppdm_metric); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_metric values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	ppdm_metric.Row_created_date = formatDate(ppdm_metric.Row_created_date)
	ppdm_metric.Row_changed_date = formatDate(ppdm_metric.Row_changed_date)
	ppdm_metric.Effective_date = formatDateString(ppdm_metric.Effective_date)
	ppdm_metric.End_date = formatDateString(ppdm_metric.End_date)
	ppdm_metric.Expiry_date = formatDateString(ppdm_metric.Expiry_date)
	ppdm_metric.Start_date = formatDateString(ppdm_metric.Start_date)
	ppdm_metric.Row_effective_date = formatDateString(ppdm_metric.Row_effective_date)
	ppdm_metric.Row_expiry_date = formatDateString(ppdm_metric.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_metric.Metric_id, ppdm_metric.Active_ind, ppdm_metric.Effective_date, ppdm_metric.End_date, ppdm_metric.Expiry_date, ppdm_metric.Metric_name, ppdm_metric.Metric_procedure, ppdm_metric.Metric_type, ppdm_metric.Organization_id, ppdm_metric.Organization_seq_no, ppdm_metric.Owner_ba_id, ppdm_metric.Ppdm_guid, ppdm_metric.Procedure_id, ppdm_metric.Procedure_system_id, ppdm_metric.Projected_final_count, ppdm_metric.Purpose_desc, ppdm_metric.Remark, ppdm_metric.Source, ppdm_metric.Start_date, ppdm_metric.Row_changed_by, ppdm_metric.Row_changed_date, ppdm_metric.Row_created_by, ppdm_metric.Row_created_date, ppdm_metric.Row_effective_date, ppdm_metric.Row_expiry_date, ppdm_metric.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmMetric(c *fiber.Ctx) error {
	var ppdm_metric dto.Ppdm_metric

	setDefaults(&ppdm_metric)

	if err := json.Unmarshal(c.Body(), &ppdm_metric); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_metric.Metric_id = id

    if ppdm_metric.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_metric set active_ind = :1, effective_date = :2, end_date = :3, expiry_date = :4, metric_name = :5, metric_procedure = :6, metric_type = :7, organization_id = :8, organization_seq_no = :9, owner_ba_id = :10, ppdm_guid = :11, procedure_id = :12, procedure_system_id = :13, projected_final_count = :14, purpose_desc = :15, remark = :16, source = :17, start_date = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where metric_id = :26")
	if err != nil {
		return err
	}

	ppdm_metric.Row_changed_date = formatDate(ppdm_metric.Row_changed_date)
	ppdm_metric.Effective_date = formatDateString(ppdm_metric.Effective_date)
	ppdm_metric.End_date = formatDateString(ppdm_metric.End_date)
	ppdm_metric.Expiry_date = formatDateString(ppdm_metric.Expiry_date)
	ppdm_metric.Start_date = formatDateString(ppdm_metric.Start_date)
	ppdm_metric.Row_effective_date = formatDateString(ppdm_metric.Row_effective_date)
	ppdm_metric.Row_expiry_date = formatDateString(ppdm_metric.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_metric.Active_ind, ppdm_metric.Effective_date, ppdm_metric.End_date, ppdm_metric.Expiry_date, ppdm_metric.Metric_name, ppdm_metric.Metric_procedure, ppdm_metric.Metric_type, ppdm_metric.Organization_id, ppdm_metric.Organization_seq_no, ppdm_metric.Owner_ba_id, ppdm_metric.Ppdm_guid, ppdm_metric.Procedure_id, ppdm_metric.Procedure_system_id, ppdm_metric.Projected_final_count, ppdm_metric.Purpose_desc, ppdm_metric.Remark, ppdm_metric.Source, ppdm_metric.Start_date, ppdm_metric.Row_changed_by, ppdm_metric.Row_changed_date, ppdm_metric.Row_created_by, ppdm_metric.Row_effective_date, ppdm_metric.Row_expiry_date, ppdm_metric.Row_quality, ppdm_metric.Metric_id)
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

func PatchPpdmMetric(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_metric set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where metric_id = :id"

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

func DeletePpdmMetric(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_metric dto.Ppdm_metric
	ppdm_metric.Metric_id = id

	stmt, err := db.Prepare("delete from ppdm_metric where metric_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_metric.Metric_id)
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


