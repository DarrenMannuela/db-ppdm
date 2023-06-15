package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoFossilList(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_fossil_list")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_fossil_list

	for rows.Next() {
		var paleo_fossil_list dto.Paleo_fossil_list
		if err := rows.Scan(&paleo_fossil_list.Paleo_summary_id, &paleo_fossil_list.Fossil_detail_id, &paleo_fossil_list.Fossil_id, &paleo_fossil_list.Abundance_count, &paleo_fossil_list.Abundance_qualifier_id, &paleo_fossil_list.Active_ind, &paleo_fossil_list.Analysis_id, &paleo_fossil_list.Anl_source, &paleo_fossil_list.Confidence_id, &paleo_fossil_list.Depth, &paleo_fossil_list.Depth_ouom, &paleo_fossil_list.Effective_date, &paleo_fossil_list.Expiry_date, &paleo_fossil_list.Fossil_color, &paleo_fossil_list.Lith_sample_id, &paleo_fossil_list.Maturation_obs_no, &paleo_fossil_list.Ontogeny_type, &paleo_fossil_list.Paleo_analyst_ba_id, &paleo_fossil_list.Physical_item_id, &paleo_fossil_list.Ppdm_guid, &paleo_fossil_list.Preferred_ind, &paleo_fossil_list.Preservation_quality, &paleo_fossil_list.Preservation_type, &paleo_fossil_list.Remark, &paleo_fossil_list.Slide_loc_x, &paleo_fossil_list.Slide_loc_y, &paleo_fossil_list.Slide_num, &paleo_fossil_list.Source, &paleo_fossil_list.Tai_color, &paleo_fossil_list.Type_fossil_type, &paleo_fossil_list.Row_changed_by, &paleo_fossil_list.Row_changed_date, &paleo_fossil_list.Row_created_by, &paleo_fossil_list.Row_created_date, &paleo_fossil_list.Row_effective_date, &paleo_fossil_list.Row_expiry_date, &paleo_fossil_list.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_fossil_list to result
		result = append(result, paleo_fossil_list)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoFossilList(c *fiber.Ctx) error {
	var paleo_fossil_list dto.Paleo_fossil_list

	setDefaults(&paleo_fossil_list)

	if err := json.Unmarshal(c.Body(), &paleo_fossil_list); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_fossil_list values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	paleo_fossil_list.Row_created_date = formatDate(paleo_fossil_list.Row_created_date)
	paleo_fossil_list.Row_changed_date = formatDate(paleo_fossil_list.Row_changed_date)
	paleo_fossil_list.Effective_date = formatDateString(paleo_fossil_list.Effective_date)
	paleo_fossil_list.Expiry_date = formatDateString(paleo_fossil_list.Expiry_date)
	paleo_fossil_list.Row_effective_date = formatDateString(paleo_fossil_list.Row_effective_date)
	paleo_fossil_list.Row_expiry_date = formatDateString(paleo_fossil_list.Row_expiry_date)






	rows, err := stmt.Exec(paleo_fossil_list.Paleo_summary_id, paleo_fossil_list.Fossil_detail_id, paleo_fossil_list.Fossil_id, paleo_fossil_list.Abundance_count, paleo_fossil_list.Abundance_qualifier_id, paleo_fossil_list.Active_ind, paleo_fossil_list.Analysis_id, paleo_fossil_list.Anl_source, paleo_fossil_list.Confidence_id, paleo_fossil_list.Depth, paleo_fossil_list.Depth_ouom, paleo_fossil_list.Effective_date, paleo_fossil_list.Expiry_date, paleo_fossil_list.Fossil_color, paleo_fossil_list.Lith_sample_id, paleo_fossil_list.Maturation_obs_no, paleo_fossil_list.Ontogeny_type, paleo_fossil_list.Paleo_analyst_ba_id, paleo_fossil_list.Physical_item_id, paleo_fossil_list.Ppdm_guid, paleo_fossil_list.Preferred_ind, paleo_fossil_list.Preservation_quality, paleo_fossil_list.Preservation_type, paleo_fossil_list.Remark, paleo_fossil_list.Slide_loc_x, paleo_fossil_list.Slide_loc_y, paleo_fossil_list.Slide_num, paleo_fossil_list.Source, paleo_fossil_list.Tai_color, paleo_fossil_list.Type_fossil_type, paleo_fossil_list.Row_changed_by, paleo_fossil_list.Row_changed_date, paleo_fossil_list.Row_created_by, paleo_fossil_list.Row_created_date, paleo_fossil_list.Row_effective_date, paleo_fossil_list.Row_expiry_date, paleo_fossil_list.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoFossilList(c *fiber.Ctx) error {
	var paleo_fossil_list dto.Paleo_fossil_list

	setDefaults(&paleo_fossil_list)

	if err := json.Unmarshal(c.Body(), &paleo_fossil_list); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_fossil_list.Paleo_summary_id = id

    if paleo_fossil_list.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_fossil_list set fossil_detail_id = :1, fossil_id = :2, abundance_count = :3, abundance_qualifier_id = :4, active_ind = :5, analysis_id = :6, anl_source = :7, confidence_id = :8, depth = :9, depth_ouom = :10, effective_date = :11, expiry_date = :12, fossil_color = :13, lith_sample_id = :14, maturation_obs_no = :15, ontogeny_type = :16, paleo_analyst_ba_id = :17, physical_item_id = :18, ppdm_guid = :19, preferred_ind = :20, preservation_quality = :21, preservation_type = :22, remark = :23, slide_loc_x = :24, slide_loc_y = :25, slide_num = :26, source = :27, tai_color = :28, type_fossil_type = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where paleo_summary_id = :37")
	if err != nil {
		return err
	}

	paleo_fossil_list.Row_changed_date = formatDate(paleo_fossil_list.Row_changed_date)
	paleo_fossil_list.Effective_date = formatDateString(paleo_fossil_list.Effective_date)
	paleo_fossil_list.Expiry_date = formatDateString(paleo_fossil_list.Expiry_date)
	paleo_fossil_list.Row_effective_date = formatDateString(paleo_fossil_list.Row_effective_date)
	paleo_fossil_list.Row_expiry_date = formatDateString(paleo_fossil_list.Row_expiry_date)






	rows, err := stmt.Exec(paleo_fossil_list.Fossil_detail_id, paleo_fossil_list.Fossil_id, paleo_fossil_list.Abundance_count, paleo_fossil_list.Abundance_qualifier_id, paleo_fossil_list.Active_ind, paleo_fossil_list.Analysis_id, paleo_fossil_list.Anl_source, paleo_fossil_list.Confidence_id, paleo_fossil_list.Depth, paleo_fossil_list.Depth_ouom, paleo_fossil_list.Effective_date, paleo_fossil_list.Expiry_date, paleo_fossil_list.Fossil_color, paleo_fossil_list.Lith_sample_id, paleo_fossil_list.Maturation_obs_no, paleo_fossil_list.Ontogeny_type, paleo_fossil_list.Paleo_analyst_ba_id, paleo_fossil_list.Physical_item_id, paleo_fossil_list.Ppdm_guid, paleo_fossil_list.Preferred_ind, paleo_fossil_list.Preservation_quality, paleo_fossil_list.Preservation_type, paleo_fossil_list.Remark, paleo_fossil_list.Slide_loc_x, paleo_fossil_list.Slide_loc_y, paleo_fossil_list.Slide_num, paleo_fossil_list.Source, paleo_fossil_list.Tai_color, paleo_fossil_list.Type_fossil_type, paleo_fossil_list.Row_changed_by, paleo_fossil_list.Row_changed_date, paleo_fossil_list.Row_created_by, paleo_fossil_list.Row_effective_date, paleo_fossil_list.Row_expiry_date, paleo_fossil_list.Row_quality, paleo_fossil_list.Paleo_summary_id)
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

func PatchPaleoFossilList(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_fossil_list set "
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
	query += " where paleo_summary_id = :id"

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

func DeletePaleoFossilList(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_fossil_list dto.Paleo_fossil_list
	paleo_fossil_list.Paleo_summary_id = id

	stmt, err := db.Prepare("delete from paleo_fossil_list where paleo_summary_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_fossil_list.Paleo_summary_id)
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


