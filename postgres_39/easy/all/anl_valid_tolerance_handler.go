package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlValidTolerance(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_valid_tolerance")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_valid_tolerance

	for rows.Next() {
		var anl_valid_tolerance dto.Anl_valid_tolerance
		if err := rows.Scan(&anl_valid_tolerance.Method_set_id, &anl_valid_tolerance.Method_id, &anl_valid_tolerance.Equip_obs_no, &anl_valid_tolerance.Tolerance_obs_no, &anl_valid_tolerance.Active_ind, &anl_valid_tolerance.Effective_date, &anl_valid_tolerance.Expiry_date, &anl_valid_tolerance.Max_tolerance, &anl_valid_tolerance.Max_tolerance_ouom, &anl_valid_tolerance.Max_tolerance_rep, &anl_valid_tolerance.Max_tolerance_uom, &anl_valid_tolerance.Min_tolerance, &anl_valid_tolerance.Min_tolerance_ouom, &anl_valid_tolerance.Min_tolerance_rep, &anl_valid_tolerance.Min_tolerance_uom, &anl_valid_tolerance.Ppdm_guid, &anl_valid_tolerance.Remark, &anl_valid_tolerance.Source, &anl_valid_tolerance.Substance_id, &anl_valid_tolerance.Tolerance_desc, &anl_valid_tolerance.Tolerance_type, &anl_valid_tolerance.Row_changed_by, &anl_valid_tolerance.Row_changed_date, &anl_valid_tolerance.Row_created_by, &anl_valid_tolerance.Row_created_date, &anl_valid_tolerance.Row_effective_date, &anl_valid_tolerance.Row_expiry_date, &anl_valid_tolerance.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_valid_tolerance to result
		result = append(result, anl_valid_tolerance)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlValidTolerance(c *fiber.Ctx) error {
	var anl_valid_tolerance dto.Anl_valid_tolerance

	setDefaults(&anl_valid_tolerance)

	if err := json.Unmarshal(c.Body(), &anl_valid_tolerance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_valid_tolerance values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	anl_valid_tolerance.Row_created_date = formatDate(anl_valid_tolerance.Row_created_date)
	anl_valid_tolerance.Row_changed_date = formatDate(anl_valid_tolerance.Row_changed_date)
	anl_valid_tolerance.Effective_date = formatDateString(anl_valid_tolerance.Effective_date)
	anl_valid_tolerance.Expiry_date = formatDateString(anl_valid_tolerance.Expiry_date)
	anl_valid_tolerance.Row_effective_date = formatDateString(anl_valid_tolerance.Row_effective_date)
	anl_valid_tolerance.Row_expiry_date = formatDateString(anl_valid_tolerance.Row_expiry_date)






	rows, err := stmt.Exec(anl_valid_tolerance.Method_set_id, anl_valid_tolerance.Method_id, anl_valid_tolerance.Equip_obs_no, anl_valid_tolerance.Tolerance_obs_no, anl_valid_tolerance.Active_ind, anl_valid_tolerance.Effective_date, anl_valid_tolerance.Expiry_date, anl_valid_tolerance.Max_tolerance, anl_valid_tolerance.Max_tolerance_ouom, anl_valid_tolerance.Max_tolerance_rep, anl_valid_tolerance.Max_tolerance_uom, anl_valid_tolerance.Min_tolerance, anl_valid_tolerance.Min_tolerance_ouom, anl_valid_tolerance.Min_tolerance_rep, anl_valid_tolerance.Min_tolerance_uom, anl_valid_tolerance.Ppdm_guid, anl_valid_tolerance.Remark, anl_valid_tolerance.Source, anl_valid_tolerance.Substance_id, anl_valid_tolerance.Tolerance_desc, anl_valid_tolerance.Tolerance_type, anl_valid_tolerance.Row_changed_by, anl_valid_tolerance.Row_changed_date, anl_valid_tolerance.Row_created_by, anl_valid_tolerance.Row_created_date, anl_valid_tolerance.Row_effective_date, anl_valid_tolerance.Row_expiry_date, anl_valid_tolerance.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlValidTolerance(c *fiber.Ctx) error {
	var anl_valid_tolerance dto.Anl_valid_tolerance

	setDefaults(&anl_valid_tolerance)

	if err := json.Unmarshal(c.Body(), &anl_valid_tolerance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_valid_tolerance.Method_set_id = id

    if anl_valid_tolerance.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_valid_tolerance set method_id = :1, equip_obs_no = :2, tolerance_obs_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, max_tolerance = :7, max_tolerance_ouom = :8, max_tolerance_rep = :9, max_tolerance_uom = :10, min_tolerance = :11, min_tolerance_ouom = :12, min_tolerance_rep = :13, min_tolerance_uom = :14, ppdm_guid = :15, remark = :16, source = :17, substance_id = :18, tolerance_desc = :19, tolerance_type = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where method_set_id = :28")
	if err != nil {
		return err
	}

	anl_valid_tolerance.Row_changed_date = formatDate(anl_valid_tolerance.Row_changed_date)
	anl_valid_tolerance.Effective_date = formatDateString(anl_valid_tolerance.Effective_date)
	anl_valid_tolerance.Expiry_date = formatDateString(anl_valid_tolerance.Expiry_date)
	anl_valid_tolerance.Row_effective_date = formatDateString(anl_valid_tolerance.Row_effective_date)
	anl_valid_tolerance.Row_expiry_date = formatDateString(anl_valid_tolerance.Row_expiry_date)






	rows, err := stmt.Exec(anl_valid_tolerance.Method_id, anl_valid_tolerance.Equip_obs_no, anl_valid_tolerance.Tolerance_obs_no, anl_valid_tolerance.Active_ind, anl_valid_tolerance.Effective_date, anl_valid_tolerance.Expiry_date, anl_valid_tolerance.Max_tolerance, anl_valid_tolerance.Max_tolerance_ouom, anl_valid_tolerance.Max_tolerance_rep, anl_valid_tolerance.Max_tolerance_uom, anl_valid_tolerance.Min_tolerance, anl_valid_tolerance.Min_tolerance_ouom, anl_valid_tolerance.Min_tolerance_rep, anl_valid_tolerance.Min_tolerance_uom, anl_valid_tolerance.Ppdm_guid, anl_valid_tolerance.Remark, anl_valid_tolerance.Source, anl_valid_tolerance.Substance_id, anl_valid_tolerance.Tolerance_desc, anl_valid_tolerance.Tolerance_type, anl_valid_tolerance.Row_changed_by, anl_valid_tolerance.Row_changed_date, anl_valid_tolerance.Row_created_by, anl_valid_tolerance.Row_effective_date, anl_valid_tolerance.Row_expiry_date, anl_valid_tolerance.Row_quality, anl_valid_tolerance.Method_set_id)
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

func PatchAnlValidTolerance(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_valid_tolerance set "
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
	query += " where method_set_id = :id"

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

func DeleteAnlValidTolerance(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_valid_tolerance dto.Anl_valid_tolerance
	anl_valid_tolerance.Method_set_id = id

	stmt, err := db.Prepare("delete from anl_valid_tolerance where method_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_valid_tolerance.Method_set_id)
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


