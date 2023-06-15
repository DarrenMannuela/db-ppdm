package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisPointFlow(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_point_flow")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_point_flow

	for rows.Next() {
		var seis_point_flow dto.Seis_point_flow
		if err := rows.Scan(&seis_point_flow.Seis_set_subtype, &seis_point_flow.Seis_set_id, &seis_point_flow.Seis_point_id, &seis_point_flow.Flow_id, &seis_point_flow.Active_ind, &seis_point_flow.Effective_date, &seis_point_flow.Expiry_date, &seis_point_flow.Flow_depth, &seis_point_flow.Flow_depth_ouom, &seis_point_flow.Flow_duration, &seis_point_flow.Flow_duration_uom, &seis_point_flow.Flow_volume, &seis_point_flow.Flow_volume_ouom, &seis_point_flow.Ppdm_guid, &seis_point_flow.Remark, &seis_point_flow.Remedial_ba_id, &seis_point_flow.Remedial_date, &seis_point_flow.Remedial_ind, &seis_point_flow.Report_date, &seis_point_flow.Source, &seis_point_flow.Row_changed_by, &seis_point_flow.Row_changed_date, &seis_point_flow.Row_created_by, &seis_point_flow.Row_created_date, &seis_point_flow.Row_effective_date, &seis_point_flow.Row_expiry_date, &seis_point_flow.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_point_flow to result
		result = append(result, seis_point_flow)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisPointFlow(c *fiber.Ctx) error {
	var seis_point_flow dto.Seis_point_flow

	setDefaults(&seis_point_flow)

	if err := json.Unmarshal(c.Body(), &seis_point_flow); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_point_flow values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	seis_point_flow.Row_created_date = formatDate(seis_point_flow.Row_created_date)
	seis_point_flow.Row_changed_date = formatDate(seis_point_flow.Row_changed_date)
	seis_point_flow.Effective_date = formatDateString(seis_point_flow.Effective_date)
	seis_point_flow.Expiry_date = formatDateString(seis_point_flow.Expiry_date)
	seis_point_flow.Remedial_date = formatDateString(seis_point_flow.Remedial_date)
	seis_point_flow.Report_date = formatDateString(seis_point_flow.Report_date)
	seis_point_flow.Row_effective_date = formatDateString(seis_point_flow.Row_effective_date)
	seis_point_flow.Row_expiry_date = formatDateString(seis_point_flow.Row_expiry_date)






	rows, err := stmt.Exec(seis_point_flow.Seis_set_subtype, seis_point_flow.Seis_set_id, seis_point_flow.Seis_point_id, seis_point_flow.Flow_id, seis_point_flow.Active_ind, seis_point_flow.Effective_date, seis_point_flow.Expiry_date, seis_point_flow.Flow_depth, seis_point_flow.Flow_depth_ouom, seis_point_flow.Flow_duration, seis_point_flow.Flow_duration_uom, seis_point_flow.Flow_volume, seis_point_flow.Flow_volume_ouom, seis_point_flow.Ppdm_guid, seis_point_flow.Remark, seis_point_flow.Remedial_ba_id, seis_point_flow.Remedial_date, seis_point_flow.Remedial_ind, seis_point_flow.Report_date, seis_point_flow.Source, seis_point_flow.Row_changed_by, seis_point_flow.Row_changed_date, seis_point_flow.Row_created_by, seis_point_flow.Row_created_date, seis_point_flow.Row_effective_date, seis_point_flow.Row_expiry_date, seis_point_flow.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisPointFlow(c *fiber.Ctx) error {
	var seis_point_flow dto.Seis_point_flow

	setDefaults(&seis_point_flow)

	if err := json.Unmarshal(c.Body(), &seis_point_flow); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_point_flow.Seis_set_subtype = id

    if seis_point_flow.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_point_flow set seis_set_id = :1, seis_point_id = :2, flow_id = :3, active_ind = :4, effective_date = :5, expiry_date = :6, flow_depth = :7, flow_depth_ouom = :8, flow_duration = :9, flow_duration_uom = :10, flow_volume = :11, flow_volume_ouom = :12, ppdm_guid = :13, remark = :14, remedial_ba_id = :15, remedial_date = :16, remedial_ind = :17, report_date = :18, source = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where seis_set_subtype = :27")
	if err != nil {
		return err
	}

	seis_point_flow.Row_changed_date = formatDate(seis_point_flow.Row_changed_date)
	seis_point_flow.Effective_date = formatDateString(seis_point_flow.Effective_date)
	seis_point_flow.Expiry_date = formatDateString(seis_point_flow.Expiry_date)
	seis_point_flow.Remedial_date = formatDateString(seis_point_flow.Remedial_date)
	seis_point_flow.Report_date = formatDateString(seis_point_flow.Report_date)
	seis_point_flow.Row_effective_date = formatDateString(seis_point_flow.Row_effective_date)
	seis_point_flow.Row_expiry_date = formatDateString(seis_point_flow.Row_expiry_date)






	rows, err := stmt.Exec(seis_point_flow.Seis_set_id, seis_point_flow.Seis_point_id, seis_point_flow.Flow_id, seis_point_flow.Active_ind, seis_point_flow.Effective_date, seis_point_flow.Expiry_date, seis_point_flow.Flow_depth, seis_point_flow.Flow_depth_ouom, seis_point_flow.Flow_duration, seis_point_flow.Flow_duration_uom, seis_point_flow.Flow_volume, seis_point_flow.Flow_volume_ouom, seis_point_flow.Ppdm_guid, seis_point_flow.Remark, seis_point_flow.Remedial_ba_id, seis_point_flow.Remedial_date, seis_point_flow.Remedial_ind, seis_point_flow.Report_date, seis_point_flow.Source, seis_point_flow.Row_changed_by, seis_point_flow.Row_changed_date, seis_point_flow.Row_created_by, seis_point_flow.Row_effective_date, seis_point_flow.Row_expiry_date, seis_point_flow.Row_quality, seis_point_flow.Seis_set_subtype)
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

func PatchSeisPointFlow(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_point_flow set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "remedial_date" || key == "report_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where seis_set_subtype = :id"

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

func DeleteSeisPointFlow(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_point_flow dto.Seis_point_flow
	seis_point_flow.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_point_flow where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_point_flow.Seis_set_subtype)
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


