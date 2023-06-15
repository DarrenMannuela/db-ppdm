package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisSetPlan(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_set_plan")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_set_plan

	for rows.Next() {
		var seis_set_plan dto.Seis_set_plan
		if err := rows.Scan(&seis_set_plan.Seis_set_subtype, &seis_set_plan.Seis_set_id, &seis_set_plan.Active_ind, &seis_set_plan.Effective_date, &seis_set_plan.Expiry_date, &seis_set_plan.Planned_end_date, &seis_set_plan.Planned_length, &seis_set_plan.Planned_length_ouom, &seis_set_plan.Planned_start_date, &seis_set_plan.Ppdm_guid, &seis_set_plan.Remark, &seis_set_plan.Source, &seis_set_plan.Row_changed_by, &seis_set_plan.Row_changed_date, &seis_set_plan.Row_created_by, &seis_set_plan.Row_created_date, &seis_set_plan.Row_effective_date, &seis_set_plan.Row_expiry_date, &seis_set_plan.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_set_plan to result
		result = append(result, seis_set_plan)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisSetPlan(c *fiber.Ctx) error {
	var seis_set_plan dto.Seis_set_plan

	setDefaults(&seis_set_plan)

	if err := json.Unmarshal(c.Body(), &seis_set_plan); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_set_plan values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	seis_set_plan.Row_created_date = formatDate(seis_set_plan.Row_created_date)
	seis_set_plan.Row_changed_date = formatDate(seis_set_plan.Row_changed_date)
	seis_set_plan.Effective_date = formatDateString(seis_set_plan.Effective_date)
	seis_set_plan.Expiry_date = formatDateString(seis_set_plan.Expiry_date)
	seis_set_plan.Planned_end_date = formatDateString(seis_set_plan.Planned_end_date)
	seis_set_plan.Planned_start_date = formatDateString(seis_set_plan.Planned_start_date)
	seis_set_plan.Row_effective_date = formatDateString(seis_set_plan.Row_effective_date)
	seis_set_plan.Row_expiry_date = formatDateString(seis_set_plan.Row_expiry_date)






	rows, err := stmt.Exec(seis_set_plan.Seis_set_subtype, seis_set_plan.Seis_set_id, seis_set_plan.Active_ind, seis_set_plan.Effective_date, seis_set_plan.Expiry_date, seis_set_plan.Planned_end_date, seis_set_plan.Planned_length, seis_set_plan.Planned_length_ouom, seis_set_plan.Planned_start_date, seis_set_plan.Ppdm_guid, seis_set_plan.Remark, seis_set_plan.Source, seis_set_plan.Row_changed_by, seis_set_plan.Row_changed_date, seis_set_plan.Row_created_by, seis_set_plan.Row_created_date, seis_set_plan.Row_effective_date, seis_set_plan.Row_expiry_date, seis_set_plan.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisSetPlan(c *fiber.Ctx) error {
	var seis_set_plan dto.Seis_set_plan

	setDefaults(&seis_set_plan)

	if err := json.Unmarshal(c.Body(), &seis_set_plan); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_set_plan.Seis_set_subtype = id

    if seis_set_plan.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_set_plan set seis_set_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, planned_end_date = :5, planned_length = :6, planned_length_ouom = :7, planned_start_date = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where seis_set_subtype = :19")
	if err != nil {
		return err
	}

	seis_set_plan.Row_changed_date = formatDate(seis_set_plan.Row_changed_date)
	seis_set_plan.Effective_date = formatDateString(seis_set_plan.Effective_date)
	seis_set_plan.Expiry_date = formatDateString(seis_set_plan.Expiry_date)
	seis_set_plan.Planned_end_date = formatDateString(seis_set_plan.Planned_end_date)
	seis_set_plan.Planned_start_date = formatDateString(seis_set_plan.Planned_start_date)
	seis_set_plan.Row_effective_date = formatDateString(seis_set_plan.Row_effective_date)
	seis_set_plan.Row_expiry_date = formatDateString(seis_set_plan.Row_expiry_date)






	rows, err := stmt.Exec(seis_set_plan.Seis_set_id, seis_set_plan.Active_ind, seis_set_plan.Effective_date, seis_set_plan.Expiry_date, seis_set_plan.Planned_end_date, seis_set_plan.Planned_length, seis_set_plan.Planned_length_ouom, seis_set_plan.Planned_start_date, seis_set_plan.Ppdm_guid, seis_set_plan.Remark, seis_set_plan.Source, seis_set_plan.Row_changed_by, seis_set_plan.Row_changed_date, seis_set_plan.Row_created_by, seis_set_plan.Row_effective_date, seis_set_plan.Row_expiry_date, seis_set_plan.Row_quality, seis_set_plan.Seis_set_subtype)
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

func PatchSeisSetPlan(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_set_plan set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "planned_end_date" || key == "planned_start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSeisSetPlan(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_set_plan dto.Seis_set_plan
	seis_set_plan.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_set_plan where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_set_plan.Seis_set_subtype)
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


