package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpDescXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_desc_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_desc_xref

	for rows.Next() {
		var sp_desc_xref dto.Sp_desc_xref
		if err := rows.Scan(&sp_desc_xref.Spatial_description_id, &sp_desc_xref.Spatial_obs_no, &sp_desc_xref.Spatial_description_id_2, &sp_desc_xref.Spatial_obs_no_2, &sp_desc_xref.Active_ind, &sp_desc_xref.Effective_date, &sp_desc_xref.Expiry_date, &sp_desc_xref.Overlap_size, &sp_desc_xref.Overlap_size_ouom, &sp_desc_xref.Ppdm_guid, &sp_desc_xref.Remark, &sp_desc_xref.Source, &sp_desc_xref.Spatial_xref_type, &sp_desc_xref.Row_changed_by, &sp_desc_xref.Row_changed_date, &sp_desc_xref.Row_created_by, &sp_desc_xref.Row_created_date, &sp_desc_xref.Row_effective_date, &sp_desc_xref.Row_expiry_date, &sp_desc_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_desc_xref to result
		result = append(result, sp_desc_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpDescXref(c *fiber.Ctx) error {
	var sp_desc_xref dto.Sp_desc_xref

	setDefaults(&sp_desc_xref)

	if err := json.Unmarshal(c.Body(), &sp_desc_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_desc_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	sp_desc_xref.Row_created_date = formatDate(sp_desc_xref.Row_created_date)
	sp_desc_xref.Row_changed_date = formatDate(sp_desc_xref.Row_changed_date)
	sp_desc_xref.Effective_date = formatDateString(sp_desc_xref.Effective_date)
	sp_desc_xref.Expiry_date = formatDateString(sp_desc_xref.Expiry_date)
	sp_desc_xref.Row_effective_date = formatDateString(sp_desc_xref.Row_effective_date)
	sp_desc_xref.Row_expiry_date = formatDateString(sp_desc_xref.Row_expiry_date)






	rows, err := stmt.Exec(sp_desc_xref.Spatial_description_id, sp_desc_xref.Spatial_obs_no, sp_desc_xref.Spatial_description_id_2, sp_desc_xref.Spatial_obs_no_2, sp_desc_xref.Active_ind, sp_desc_xref.Effective_date, sp_desc_xref.Expiry_date, sp_desc_xref.Overlap_size, sp_desc_xref.Overlap_size_ouom, sp_desc_xref.Ppdm_guid, sp_desc_xref.Remark, sp_desc_xref.Source, sp_desc_xref.Spatial_xref_type, sp_desc_xref.Row_changed_by, sp_desc_xref.Row_changed_date, sp_desc_xref.Row_created_by, sp_desc_xref.Row_created_date, sp_desc_xref.Row_effective_date, sp_desc_xref.Row_expiry_date, sp_desc_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpDescXref(c *fiber.Ctx) error {
	var sp_desc_xref dto.Sp_desc_xref

	setDefaults(&sp_desc_xref)

	if err := json.Unmarshal(c.Body(), &sp_desc_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_desc_xref.Spatial_description_id = id

    if sp_desc_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_desc_xref set spatial_obs_no = :1, spatial_description_id_2 = :2, spatial_obs_no_2 = :3, active_ind = :4, effective_date = :5, expiry_date = :6, overlap_size = :7, overlap_size_ouom = :8, ppdm_guid = :9, remark = :10, source = :11, spatial_xref_type = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where spatial_description_id = :20")
	if err != nil {
		return err
	}

	sp_desc_xref.Row_changed_date = formatDate(sp_desc_xref.Row_changed_date)
	sp_desc_xref.Effective_date = formatDateString(sp_desc_xref.Effective_date)
	sp_desc_xref.Expiry_date = formatDateString(sp_desc_xref.Expiry_date)
	sp_desc_xref.Row_effective_date = formatDateString(sp_desc_xref.Row_effective_date)
	sp_desc_xref.Row_expiry_date = formatDateString(sp_desc_xref.Row_expiry_date)






	rows, err := stmt.Exec(sp_desc_xref.Spatial_obs_no, sp_desc_xref.Spatial_description_id_2, sp_desc_xref.Spatial_obs_no_2, sp_desc_xref.Active_ind, sp_desc_xref.Effective_date, sp_desc_xref.Expiry_date, sp_desc_xref.Overlap_size, sp_desc_xref.Overlap_size_ouom, sp_desc_xref.Ppdm_guid, sp_desc_xref.Remark, sp_desc_xref.Source, sp_desc_xref.Spatial_xref_type, sp_desc_xref.Row_changed_by, sp_desc_xref.Row_changed_date, sp_desc_xref.Row_created_by, sp_desc_xref.Row_effective_date, sp_desc_xref.Row_expiry_date, sp_desc_xref.Row_quality, sp_desc_xref.Spatial_description_id)
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

func PatchSpDescXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_desc_xref set "
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

func DeleteSpDescXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_desc_xref dto.Sp_desc_xref
	sp_desc_xref.Spatial_description_id = id

	stmt, err := db.Prepare("delete from sp_desc_xref where spatial_description_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_desc_xref.Spatial_description_id)
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


