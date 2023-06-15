package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlPaleoMaturity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_paleo_maturity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_paleo_maturity

	for rows.Next() {
		var anl_paleo_maturity dto.Anl_paleo_maturity
		if err := rows.Scan(&anl_paleo_maturity.Analysis_id, &anl_paleo_maturity.Anl_source, &anl_paleo_maturity.Maturation_obs_no, &anl_paleo_maturity.Active_ind, &anl_paleo_maturity.Effective_date, &anl_paleo_maturity.Expiry_date, &anl_paleo_maturity.Fluor_color, &anl_paleo_maturity.Fluor_intensity_desc, &anl_paleo_maturity.Fluor_intensity_value, &anl_paleo_maturity.Fluor_intensity_value_ouom, &anl_paleo_maturity.Fluor_intensity_value_uom, &anl_paleo_maturity.Maturity_method_type, &anl_paleo_maturity.Ppdm_guid, &anl_paleo_maturity.Preferred_ind, &anl_paleo_maturity.Problem_ind, &anl_paleo_maturity.Remark, &anl_paleo_maturity.Sci_color, &anl_paleo_maturity.Sci_max, &anl_paleo_maturity.Sci_max_ouom, &anl_paleo_maturity.Sci_max_uom, &anl_paleo_maturity.Sci_min, &anl_paleo_maturity.Sci_min_ouom, &anl_paleo_maturity.Sci_min_uom, &anl_paleo_maturity.Source, &anl_paleo_maturity.Step_seq_no, &anl_paleo_maturity.Substance_id, &anl_paleo_maturity.Tai_color, &anl_paleo_maturity.Tai_max, &anl_paleo_maturity.Tai_max_ouom, &anl_paleo_maturity.Tai_max_uom, &anl_paleo_maturity.Tai_min, &anl_paleo_maturity.Tai_min_ouom, &anl_paleo_maturity.Tai_min_uom, &anl_paleo_maturity.Row_changed_by, &anl_paleo_maturity.Row_changed_date, &anl_paleo_maturity.Row_created_by, &anl_paleo_maturity.Row_created_date, &anl_paleo_maturity.Row_effective_date, &anl_paleo_maturity.Row_expiry_date, &anl_paleo_maturity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_paleo_maturity to result
		result = append(result, anl_paleo_maturity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlPaleoMaturity(c *fiber.Ctx) error {
	var anl_paleo_maturity dto.Anl_paleo_maturity

	setDefaults(&anl_paleo_maturity)

	if err := json.Unmarshal(c.Body(), &anl_paleo_maturity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_paleo_maturity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40)")
	if err != nil {
		return err
	}
	anl_paleo_maturity.Row_created_date = formatDate(anl_paleo_maturity.Row_created_date)
	anl_paleo_maturity.Row_changed_date = formatDate(anl_paleo_maturity.Row_changed_date)
	anl_paleo_maturity.Effective_date = formatDateString(anl_paleo_maturity.Effective_date)
	anl_paleo_maturity.Expiry_date = formatDateString(anl_paleo_maturity.Expiry_date)
	anl_paleo_maturity.Row_effective_date = formatDateString(anl_paleo_maturity.Row_effective_date)
	anl_paleo_maturity.Row_expiry_date = formatDateString(anl_paleo_maturity.Row_expiry_date)






	rows, err := stmt.Exec(anl_paleo_maturity.Analysis_id, anl_paleo_maturity.Anl_source, anl_paleo_maturity.Maturation_obs_no, anl_paleo_maturity.Active_ind, anl_paleo_maturity.Effective_date, anl_paleo_maturity.Expiry_date, anl_paleo_maturity.Fluor_color, anl_paleo_maturity.Fluor_intensity_desc, anl_paleo_maturity.Fluor_intensity_value, anl_paleo_maturity.Fluor_intensity_value_ouom, anl_paleo_maturity.Fluor_intensity_value_uom, anl_paleo_maturity.Maturity_method_type, anl_paleo_maturity.Ppdm_guid, anl_paleo_maturity.Preferred_ind, anl_paleo_maturity.Problem_ind, anl_paleo_maturity.Remark, anl_paleo_maturity.Sci_color, anl_paleo_maturity.Sci_max, anl_paleo_maturity.Sci_max_ouom, anl_paleo_maturity.Sci_max_uom, anl_paleo_maturity.Sci_min, anl_paleo_maturity.Sci_min_ouom, anl_paleo_maturity.Sci_min_uom, anl_paleo_maturity.Source, anl_paleo_maturity.Step_seq_no, anl_paleo_maturity.Substance_id, anl_paleo_maturity.Tai_color, anl_paleo_maturity.Tai_max, anl_paleo_maturity.Tai_max_ouom, anl_paleo_maturity.Tai_max_uom, anl_paleo_maturity.Tai_min, anl_paleo_maturity.Tai_min_ouom, anl_paleo_maturity.Tai_min_uom, anl_paleo_maturity.Row_changed_by, anl_paleo_maturity.Row_changed_date, anl_paleo_maturity.Row_created_by, anl_paleo_maturity.Row_created_date, anl_paleo_maturity.Row_effective_date, anl_paleo_maturity.Row_expiry_date, anl_paleo_maturity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlPaleoMaturity(c *fiber.Ctx) error {
	var anl_paleo_maturity dto.Anl_paleo_maturity

	setDefaults(&anl_paleo_maturity)

	if err := json.Unmarshal(c.Body(), &anl_paleo_maturity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_paleo_maturity.Analysis_id = id

    if anl_paleo_maturity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_paleo_maturity set anl_source = :1, maturation_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, fluor_color = :6, fluor_intensity_desc = :7, fluor_intensity_value = :8, fluor_intensity_value_ouom = :9, fluor_intensity_value_uom = :10, maturity_method_type = :11, ppdm_guid = :12, preferred_ind = :13, problem_ind = :14, remark = :15, sci_color = :16, sci_max = :17, sci_max_ouom = :18, sci_max_uom = :19, sci_min = :20, sci_min_ouom = :21, sci_min_uom = :22, source = :23, step_seq_no = :24, substance_id = :25, tai_color = :26, tai_max = :27, tai_max_ouom = :28, tai_max_uom = :29, tai_min = :30, tai_min_ouom = :31, tai_min_uom = :32, row_changed_by = :33, row_changed_date = :34, row_created_by = :35, row_effective_date = :36, row_expiry_date = :37, row_quality = :38 where analysis_id = :40")
	if err != nil {
		return err
	}

	anl_paleo_maturity.Row_changed_date = formatDate(anl_paleo_maturity.Row_changed_date)
	anl_paleo_maturity.Effective_date = formatDateString(anl_paleo_maturity.Effective_date)
	anl_paleo_maturity.Expiry_date = formatDateString(anl_paleo_maturity.Expiry_date)
	anl_paleo_maturity.Row_effective_date = formatDateString(anl_paleo_maturity.Row_effective_date)
	anl_paleo_maturity.Row_expiry_date = formatDateString(anl_paleo_maturity.Row_expiry_date)






	rows, err := stmt.Exec(anl_paleo_maturity.Anl_source, anl_paleo_maturity.Maturation_obs_no, anl_paleo_maturity.Active_ind, anl_paleo_maturity.Effective_date, anl_paleo_maturity.Expiry_date, anl_paleo_maturity.Fluor_color, anl_paleo_maturity.Fluor_intensity_desc, anl_paleo_maturity.Fluor_intensity_value, anl_paleo_maturity.Fluor_intensity_value_ouom, anl_paleo_maturity.Fluor_intensity_value_uom, anl_paleo_maturity.Maturity_method_type, anl_paleo_maturity.Ppdm_guid, anl_paleo_maturity.Preferred_ind, anl_paleo_maturity.Problem_ind, anl_paleo_maturity.Remark, anl_paleo_maturity.Sci_color, anl_paleo_maturity.Sci_max, anl_paleo_maturity.Sci_max_ouom, anl_paleo_maturity.Sci_max_uom, anl_paleo_maturity.Sci_min, anl_paleo_maturity.Sci_min_ouom, anl_paleo_maturity.Sci_min_uom, anl_paleo_maturity.Source, anl_paleo_maturity.Step_seq_no, anl_paleo_maturity.Substance_id, anl_paleo_maturity.Tai_color, anl_paleo_maturity.Tai_max, anl_paleo_maturity.Tai_max_ouom, anl_paleo_maturity.Tai_max_uom, anl_paleo_maturity.Tai_min, anl_paleo_maturity.Tai_min_ouom, anl_paleo_maturity.Tai_min_uom, anl_paleo_maturity.Row_changed_by, anl_paleo_maturity.Row_changed_date, anl_paleo_maturity.Row_created_by, anl_paleo_maturity.Row_effective_date, anl_paleo_maturity.Row_expiry_date, anl_paleo_maturity.Row_quality, anl_paleo_maturity.Analysis_id)
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

func PatchAnlPaleoMaturity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_paleo_maturity set "
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

func DeleteAnlPaleoMaturity(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_paleo_maturity dto.Anl_paleo_maturity
	anl_paleo_maturity.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_paleo_maturity where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_paleo_maturity.Analysis_id)
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


