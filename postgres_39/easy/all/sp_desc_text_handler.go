package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpDescText(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_desc_text")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_desc_text

	for rows.Next() {
		var sp_desc_text dto.Sp_desc_text
		if err := rows.Scan(&sp_desc_text.Spatial_description_id, &sp_desc_text.Spatial_obs_no, &sp_desc_text.Text_id, &sp_desc_text.Description_seq_no, &sp_desc_text.Active_ind, &sp_desc_text.Description, &sp_desc_text.Effective_date, &sp_desc_text.Expiry_date, &sp_desc_text.Ppdm_guid, &sp_desc_text.Remark, &sp_desc_text.Source, &sp_desc_text.Row_changed_by, &sp_desc_text.Row_changed_date, &sp_desc_text.Row_created_by, &sp_desc_text.Row_created_date, &sp_desc_text.Row_effective_date, &sp_desc_text.Row_expiry_date, &sp_desc_text.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_desc_text to result
		result = append(result, sp_desc_text)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpDescText(c *fiber.Ctx) error {
	var sp_desc_text dto.Sp_desc_text

	setDefaults(&sp_desc_text)

	if err := json.Unmarshal(c.Body(), &sp_desc_text); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_desc_text values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	sp_desc_text.Row_created_date = formatDate(sp_desc_text.Row_created_date)
	sp_desc_text.Row_changed_date = formatDate(sp_desc_text.Row_changed_date)
	sp_desc_text.Effective_date = formatDateString(sp_desc_text.Effective_date)
	sp_desc_text.Expiry_date = formatDateString(sp_desc_text.Expiry_date)
	sp_desc_text.Row_effective_date = formatDateString(sp_desc_text.Row_effective_date)
	sp_desc_text.Row_expiry_date = formatDateString(sp_desc_text.Row_expiry_date)






	rows, err := stmt.Exec(sp_desc_text.Spatial_description_id, sp_desc_text.Spatial_obs_no, sp_desc_text.Text_id, sp_desc_text.Description_seq_no, sp_desc_text.Active_ind, sp_desc_text.Description, sp_desc_text.Effective_date, sp_desc_text.Expiry_date, sp_desc_text.Ppdm_guid, sp_desc_text.Remark, sp_desc_text.Source, sp_desc_text.Row_changed_by, sp_desc_text.Row_changed_date, sp_desc_text.Row_created_by, sp_desc_text.Row_created_date, sp_desc_text.Row_effective_date, sp_desc_text.Row_expiry_date, sp_desc_text.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpDescText(c *fiber.Ctx) error {
	var sp_desc_text dto.Sp_desc_text

	setDefaults(&sp_desc_text)

	if err := json.Unmarshal(c.Body(), &sp_desc_text); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_desc_text.Spatial_description_id = id

    if sp_desc_text.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_desc_text set spatial_obs_no = :1, text_id = :2, description_seq_no = :3, active_ind = :4, description = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where spatial_description_id = :18")
	if err != nil {
		return err
	}

	sp_desc_text.Row_changed_date = formatDate(sp_desc_text.Row_changed_date)
	sp_desc_text.Effective_date = formatDateString(sp_desc_text.Effective_date)
	sp_desc_text.Expiry_date = formatDateString(sp_desc_text.Expiry_date)
	sp_desc_text.Row_effective_date = formatDateString(sp_desc_text.Row_effective_date)
	sp_desc_text.Row_expiry_date = formatDateString(sp_desc_text.Row_expiry_date)






	rows, err := stmt.Exec(sp_desc_text.Spatial_obs_no, sp_desc_text.Text_id, sp_desc_text.Description_seq_no, sp_desc_text.Active_ind, sp_desc_text.Description, sp_desc_text.Effective_date, sp_desc_text.Expiry_date, sp_desc_text.Ppdm_guid, sp_desc_text.Remark, sp_desc_text.Source, sp_desc_text.Row_changed_by, sp_desc_text.Row_changed_date, sp_desc_text.Row_created_by, sp_desc_text.Row_effective_date, sp_desc_text.Row_expiry_date, sp_desc_text.Row_quality, sp_desc_text.Spatial_description_id)
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

func PatchSpDescText(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_desc_text set "
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
	query += " where spatial_description_id = :id"

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

func DeleteSpDescText(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_desc_text dto.Sp_desc_text
	sp_desc_text.Spatial_description_id = id

	stmt, err := db.Prepare("delete from sp_desc_text where spatial_description_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_desc_text.Spatial_description_id)
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


