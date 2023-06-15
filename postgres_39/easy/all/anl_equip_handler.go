package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlEquip(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_equip")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_equip

	for rows.Next() {
		var anl_equip dto.Anl_equip
		if err := rows.Scan(&anl_equip.Analysis_id, &anl_equip.Anl_source, &anl_equip.Equip_obs_no, &anl_equip.Active_ind, &anl_equip.Callibration_record, &anl_equip.Catalogue_equip_id, &anl_equip.Effective_date, &anl_equip.Equip_id, &anl_equip.Equip_role, &anl_equip.Expiry_date, &anl_equip.Lab_ba_id, &anl_equip.Ppdm_guid, &anl_equip.Problem_ind, &anl_equip.Remark, &anl_equip.Source, &anl_equip.Step_seq_no, &anl_equip.Row_changed_by, &anl_equip.Row_changed_date, &anl_equip.Row_created_by, &anl_equip.Row_created_date, &anl_equip.Row_effective_date, &anl_equip.Row_expiry_date, &anl_equip.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_equip to result
		result = append(result, anl_equip)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlEquip(c *fiber.Ctx) error {
	var anl_equip dto.Anl_equip

	setDefaults(&anl_equip)

	if err := json.Unmarshal(c.Body(), &anl_equip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_equip values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	anl_equip.Row_created_date = formatDate(anl_equip.Row_created_date)
	anl_equip.Row_changed_date = formatDate(anl_equip.Row_changed_date)
	anl_equip.Effective_date = formatDateString(anl_equip.Effective_date)
	anl_equip.Expiry_date = formatDateString(anl_equip.Expiry_date)
	anl_equip.Row_effective_date = formatDateString(anl_equip.Row_effective_date)
	anl_equip.Row_expiry_date = formatDateString(anl_equip.Row_expiry_date)






	rows, err := stmt.Exec(anl_equip.Analysis_id, anl_equip.Anl_source, anl_equip.Equip_obs_no, anl_equip.Active_ind, anl_equip.Callibration_record, anl_equip.Catalogue_equip_id, anl_equip.Effective_date, anl_equip.Equip_id, anl_equip.Equip_role, anl_equip.Expiry_date, anl_equip.Lab_ba_id, anl_equip.Ppdm_guid, anl_equip.Problem_ind, anl_equip.Remark, anl_equip.Source, anl_equip.Step_seq_no, anl_equip.Row_changed_by, anl_equip.Row_changed_date, anl_equip.Row_created_by, anl_equip.Row_created_date, anl_equip.Row_effective_date, anl_equip.Row_expiry_date, anl_equip.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlEquip(c *fiber.Ctx) error {
	var anl_equip dto.Anl_equip

	setDefaults(&anl_equip)

	if err := json.Unmarshal(c.Body(), &anl_equip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_equip.Analysis_id = id

    if anl_equip.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_equip set anl_source = :1, equip_obs_no = :2, active_ind = :3, callibration_record = :4, catalogue_equip_id = :5, effective_date = :6, equip_id = :7, equip_role = :8, expiry_date = :9, lab_ba_id = :10, ppdm_guid = :11, problem_ind = :12, remark = :13, source = :14, step_seq_no = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where analysis_id = :23")
	if err != nil {
		return err
	}

	anl_equip.Row_changed_date = formatDate(anl_equip.Row_changed_date)
	anl_equip.Effective_date = formatDateString(anl_equip.Effective_date)
	anl_equip.Expiry_date = formatDateString(anl_equip.Expiry_date)
	anl_equip.Row_effective_date = formatDateString(anl_equip.Row_effective_date)
	anl_equip.Row_expiry_date = formatDateString(anl_equip.Row_expiry_date)






	rows, err := stmt.Exec(anl_equip.Anl_source, anl_equip.Equip_obs_no, anl_equip.Active_ind, anl_equip.Callibration_record, anl_equip.Catalogue_equip_id, anl_equip.Effective_date, anl_equip.Equip_id, anl_equip.Equip_role, anl_equip.Expiry_date, anl_equip.Lab_ba_id, anl_equip.Ppdm_guid, anl_equip.Problem_ind, anl_equip.Remark, anl_equip.Source, anl_equip.Step_seq_no, anl_equip.Row_changed_by, anl_equip.Row_changed_date, anl_equip.Row_created_by, anl_equip.Row_effective_date, anl_equip.Row_expiry_date, anl_equip.Row_quality, anl_equip.Analysis_id)
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

func PatchAnlEquip(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_equip set "
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
	query += " where analysis_id = :id"

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

func DeleteAnlEquip(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_equip dto.Anl_equip
	anl_equip.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_equip where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_equip.Analysis_id)
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


