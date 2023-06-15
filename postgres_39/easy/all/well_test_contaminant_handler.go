package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestContaminant(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_contaminant")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_contaminant

	for rows.Next() {
		var well_test_contaminant dto.Well_test_contaminant
		if err := rows.Scan(&well_test_contaminant.Uwi, &well_test_contaminant.Source, &well_test_contaminant.Test_type, &well_test_contaminant.Run_num, &well_test_contaminant.Test_num, &well_test_contaminant.Recovery_obs_no, &well_test_contaminant.Contaminant_seq_no, &well_test_contaminant.Active_ind, &well_test_contaminant.Contaminant_type, &well_test_contaminant.Effective_date, &well_test_contaminant.Expiry_date, &well_test_contaminant.Ppdm_guid, &well_test_contaminant.Remark, &well_test_contaminant.Row_changed_by, &well_test_contaminant.Row_changed_date, &well_test_contaminant.Row_created_by, &well_test_contaminant.Row_created_date, &well_test_contaminant.Row_effective_date, &well_test_contaminant.Row_expiry_date, &well_test_contaminant.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_contaminant to result
		result = append(result, well_test_contaminant)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestContaminant(c *fiber.Ctx) error {
	var well_test_contaminant dto.Well_test_contaminant

	setDefaults(&well_test_contaminant)

	if err := json.Unmarshal(c.Body(), &well_test_contaminant); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_contaminant values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	well_test_contaminant.Row_created_date = formatDate(well_test_contaminant.Row_created_date)
	well_test_contaminant.Row_changed_date = formatDate(well_test_contaminant.Row_changed_date)
	well_test_contaminant.Effective_date = formatDateString(well_test_contaminant.Effective_date)
	well_test_contaminant.Expiry_date = formatDateString(well_test_contaminant.Expiry_date)
	well_test_contaminant.Row_effective_date = formatDateString(well_test_contaminant.Row_effective_date)
	well_test_contaminant.Row_expiry_date = formatDateString(well_test_contaminant.Row_expiry_date)






	rows, err := stmt.Exec(well_test_contaminant.Uwi, well_test_contaminant.Source, well_test_contaminant.Test_type, well_test_contaminant.Run_num, well_test_contaminant.Test_num, well_test_contaminant.Recovery_obs_no, well_test_contaminant.Contaminant_seq_no, well_test_contaminant.Active_ind, well_test_contaminant.Contaminant_type, well_test_contaminant.Effective_date, well_test_contaminant.Expiry_date, well_test_contaminant.Ppdm_guid, well_test_contaminant.Remark, well_test_contaminant.Row_changed_by, well_test_contaminant.Row_changed_date, well_test_contaminant.Row_created_by, well_test_contaminant.Row_created_date, well_test_contaminant.Row_effective_date, well_test_contaminant.Row_expiry_date, well_test_contaminant.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestContaminant(c *fiber.Ctx) error {
	var well_test_contaminant dto.Well_test_contaminant

	setDefaults(&well_test_contaminant)

	if err := json.Unmarshal(c.Body(), &well_test_contaminant); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_contaminant.Uwi = id

    if well_test_contaminant.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_contaminant set source = :1, test_type = :2, run_num = :3, test_num = :4, recovery_obs_no = :5, contaminant_seq_no = :6, active_ind = :7, contaminant_type = :8, effective_date = :9, expiry_date = :10, ppdm_guid = :11, remark = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where uwi = :20")
	if err != nil {
		return err
	}

	well_test_contaminant.Row_changed_date = formatDate(well_test_contaminant.Row_changed_date)
	well_test_contaminant.Effective_date = formatDateString(well_test_contaminant.Effective_date)
	well_test_contaminant.Expiry_date = formatDateString(well_test_contaminant.Expiry_date)
	well_test_contaminant.Row_effective_date = formatDateString(well_test_contaminant.Row_effective_date)
	well_test_contaminant.Row_expiry_date = formatDateString(well_test_contaminant.Row_expiry_date)






	rows, err := stmt.Exec(well_test_contaminant.Source, well_test_contaminant.Test_type, well_test_contaminant.Run_num, well_test_contaminant.Test_num, well_test_contaminant.Recovery_obs_no, well_test_contaminant.Contaminant_seq_no, well_test_contaminant.Active_ind, well_test_contaminant.Contaminant_type, well_test_contaminant.Effective_date, well_test_contaminant.Expiry_date, well_test_contaminant.Ppdm_guid, well_test_contaminant.Remark, well_test_contaminant.Row_changed_by, well_test_contaminant.Row_changed_date, well_test_contaminant.Row_created_by, well_test_contaminant.Row_effective_date, well_test_contaminant.Row_expiry_date, well_test_contaminant.Row_quality, well_test_contaminant.Uwi)
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

func PatchWellTestContaminant(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_contaminant set "
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
	query += " where uwi = :id"

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

func DeleteWellTestContaminant(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_contaminant dto.Well_test_contaminant
	well_test_contaminant.Uwi = id

	stmt, err := db.Prepare("delete from well_test_contaminant where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_contaminant.Uwi)
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


