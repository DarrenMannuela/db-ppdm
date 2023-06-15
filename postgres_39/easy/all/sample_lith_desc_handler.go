package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSampleLithDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sample_lith_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sample_lith_desc

	for rows.Next() {
		var sample_lith_desc dto.Sample_lith_desc
		if err := rows.Scan(&sample_lith_desc.Sample_id, &sample_lith_desc.Description_id, &sample_lith_desc.Active_ind, &sample_lith_desc.Color_intensity, &sample_lith_desc.Color_type, &sample_lith_desc.Contaminant, &sample_lith_desc.Depositional_environment, &sample_lith_desc.Diagenesis_type, &sample_lith_desc.Ecozone_id, &sample_lith_desc.Effective_date, &sample_lith_desc.Expiry_date, &sample_lith_desc.Lithology, &sample_lith_desc.Oil_stain, &sample_lith_desc.Percent, &sample_lith_desc.Permeability_type, &sample_lith_desc.Porosity_type, &sample_lith_desc.Ppdm_guid, &sample_lith_desc.Remark, &sample_lith_desc.Rock_matrix, &sample_lith_desc.Rock_type, &sample_lith_desc.Source, &sample_lith_desc.Strat_name_set_id, &sample_lith_desc.Strat_unit_id, &sample_lith_desc.Row_changed_by, &sample_lith_desc.Row_changed_date, &sample_lith_desc.Row_created_by, &sample_lith_desc.Row_created_date, &sample_lith_desc.Row_effective_date, &sample_lith_desc.Row_expiry_date, &sample_lith_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sample_lith_desc to result
		result = append(result, sample_lith_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSampleLithDesc(c *fiber.Ctx) error {
	var sample_lith_desc dto.Sample_lith_desc

	setDefaults(&sample_lith_desc)

	if err := json.Unmarshal(c.Body(), &sample_lith_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sample_lith_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	sample_lith_desc.Row_created_date = formatDate(sample_lith_desc.Row_created_date)
	sample_lith_desc.Row_changed_date = formatDate(sample_lith_desc.Row_changed_date)
	sample_lith_desc.Effective_date = formatDateString(sample_lith_desc.Effective_date)
	sample_lith_desc.Expiry_date = formatDateString(sample_lith_desc.Expiry_date)
	sample_lith_desc.Row_effective_date = formatDateString(sample_lith_desc.Row_effective_date)
	sample_lith_desc.Row_expiry_date = formatDateString(sample_lith_desc.Row_expiry_date)






	rows, err := stmt.Exec(sample_lith_desc.Sample_id, sample_lith_desc.Description_id, sample_lith_desc.Active_ind, sample_lith_desc.Color_intensity, sample_lith_desc.Color_type, sample_lith_desc.Contaminant, sample_lith_desc.Depositional_environment, sample_lith_desc.Diagenesis_type, sample_lith_desc.Ecozone_id, sample_lith_desc.Effective_date, sample_lith_desc.Expiry_date, sample_lith_desc.Lithology, sample_lith_desc.Oil_stain, sample_lith_desc.Percent, sample_lith_desc.Permeability_type, sample_lith_desc.Porosity_type, sample_lith_desc.Ppdm_guid, sample_lith_desc.Remark, sample_lith_desc.Rock_matrix, sample_lith_desc.Rock_type, sample_lith_desc.Source, sample_lith_desc.Strat_name_set_id, sample_lith_desc.Strat_unit_id, sample_lith_desc.Row_changed_by, sample_lith_desc.Row_changed_date, sample_lith_desc.Row_created_by, sample_lith_desc.Row_created_date, sample_lith_desc.Row_effective_date, sample_lith_desc.Row_expiry_date, sample_lith_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSampleLithDesc(c *fiber.Ctx) error {
	var sample_lith_desc dto.Sample_lith_desc

	setDefaults(&sample_lith_desc)

	if err := json.Unmarshal(c.Body(), &sample_lith_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sample_lith_desc.Sample_id = id

    if sample_lith_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sample_lith_desc set description_id = :1, active_ind = :2, color_intensity = :3, color_type = :4, contaminant = :5, depositional_environment = :6, diagenesis_type = :7, ecozone_id = :8, effective_date = :9, expiry_date = :10, lithology = :11, oil_stain = :12, percent = :13, permeability_type = :14, porosity_type = :15, ppdm_guid = :16, remark = :17, rock_matrix = :18, rock_type = :19, source = :20, strat_name_set_id = :21, strat_unit_id = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where sample_id = :30")
	if err != nil {
		return err
	}

	sample_lith_desc.Row_changed_date = formatDate(sample_lith_desc.Row_changed_date)
	sample_lith_desc.Effective_date = formatDateString(sample_lith_desc.Effective_date)
	sample_lith_desc.Expiry_date = formatDateString(sample_lith_desc.Expiry_date)
	sample_lith_desc.Row_effective_date = formatDateString(sample_lith_desc.Row_effective_date)
	sample_lith_desc.Row_expiry_date = formatDateString(sample_lith_desc.Row_expiry_date)






	rows, err := stmt.Exec(sample_lith_desc.Description_id, sample_lith_desc.Active_ind, sample_lith_desc.Color_intensity, sample_lith_desc.Color_type, sample_lith_desc.Contaminant, sample_lith_desc.Depositional_environment, sample_lith_desc.Diagenesis_type, sample_lith_desc.Ecozone_id, sample_lith_desc.Effective_date, sample_lith_desc.Expiry_date, sample_lith_desc.Lithology, sample_lith_desc.Oil_stain, sample_lith_desc.Percent, sample_lith_desc.Permeability_type, sample_lith_desc.Porosity_type, sample_lith_desc.Ppdm_guid, sample_lith_desc.Remark, sample_lith_desc.Rock_matrix, sample_lith_desc.Rock_type, sample_lith_desc.Source, sample_lith_desc.Strat_name_set_id, sample_lith_desc.Strat_unit_id, sample_lith_desc.Row_changed_by, sample_lith_desc.Row_changed_date, sample_lith_desc.Row_created_by, sample_lith_desc.Row_effective_date, sample_lith_desc.Row_expiry_date, sample_lith_desc.Row_quality, sample_lith_desc.Sample_id)
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

func PatchSampleLithDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sample_lith_desc set "
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
	query += " where sample_id = :id"

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

func DeleteSampleLithDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var sample_lith_desc dto.Sample_lith_desc
	sample_lith_desc.Sample_id = id

	stmt, err := db.Prepare("delete from sample_lith_desc where sample_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sample_lith_desc.Sample_id)
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


