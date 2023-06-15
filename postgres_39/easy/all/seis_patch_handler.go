package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisPatch(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_patch")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_patch

	for rows.Next() {
		var seis_patch dto.Seis_patch
		if err := rows.Scan(&seis_patch.Patch_id, &seis_patch.Active_ind, &seis_patch.Channel_count, &seis_patch.Effective_date, &seis_patch.Expiry_date, &seis_patch.Gap_count, &seis_patch.Patch_layout, &seis_patch.Patch_type, &seis_patch.Ppdm_guid, &seis_patch.Remark, &seis_patch.Roll_along_method, &seis_patch.Shot_gap_ind, &seis_patch.Source, &seis_patch.Symmetric_ind, &seis_patch.Row_changed_by, &seis_patch.Row_changed_date, &seis_patch.Row_created_by, &seis_patch.Row_created_date, &seis_patch.Row_effective_date, &seis_patch.Row_expiry_date, &seis_patch.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_patch to result
		result = append(result, seis_patch)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisPatch(c *fiber.Ctx) error {
	var seis_patch dto.Seis_patch

	setDefaults(&seis_patch)

	if err := json.Unmarshal(c.Body(), &seis_patch); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_patch values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	seis_patch.Row_created_date = formatDate(seis_patch.Row_created_date)
	seis_patch.Row_changed_date = formatDate(seis_patch.Row_changed_date)
	seis_patch.Effective_date = formatDateString(seis_patch.Effective_date)
	seis_patch.Expiry_date = formatDateString(seis_patch.Expiry_date)
	seis_patch.Row_effective_date = formatDateString(seis_patch.Row_effective_date)
	seis_patch.Row_expiry_date = formatDateString(seis_patch.Row_expiry_date)






	rows, err := stmt.Exec(seis_patch.Patch_id, seis_patch.Active_ind, seis_patch.Channel_count, seis_patch.Effective_date, seis_patch.Expiry_date, seis_patch.Gap_count, seis_patch.Patch_layout, seis_patch.Patch_type, seis_patch.Ppdm_guid, seis_patch.Remark, seis_patch.Roll_along_method, seis_patch.Shot_gap_ind, seis_patch.Source, seis_patch.Symmetric_ind, seis_patch.Row_changed_by, seis_patch.Row_changed_date, seis_patch.Row_created_by, seis_patch.Row_created_date, seis_patch.Row_effective_date, seis_patch.Row_expiry_date, seis_patch.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisPatch(c *fiber.Ctx) error {
	var seis_patch dto.Seis_patch

	setDefaults(&seis_patch)

	if err := json.Unmarshal(c.Body(), &seis_patch); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_patch.Patch_id = id

    if seis_patch.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_patch set active_ind = :1, channel_count = :2, effective_date = :3, expiry_date = :4, gap_count = :5, patch_layout = :6, patch_type = :7, ppdm_guid = :8, remark = :9, roll_along_method = :10, shot_gap_ind = :11, source = :12, symmetric_ind = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where patch_id = :21")
	if err != nil {
		return err
	}

	seis_patch.Row_changed_date = formatDate(seis_patch.Row_changed_date)
	seis_patch.Effective_date = formatDateString(seis_patch.Effective_date)
	seis_patch.Expiry_date = formatDateString(seis_patch.Expiry_date)
	seis_patch.Row_effective_date = formatDateString(seis_patch.Row_effective_date)
	seis_patch.Row_expiry_date = formatDateString(seis_patch.Row_expiry_date)






	rows, err := stmt.Exec(seis_patch.Active_ind, seis_patch.Channel_count, seis_patch.Effective_date, seis_patch.Expiry_date, seis_patch.Gap_count, seis_patch.Patch_layout, seis_patch.Patch_type, seis_patch.Ppdm_guid, seis_patch.Remark, seis_patch.Roll_along_method, seis_patch.Shot_gap_ind, seis_patch.Source, seis_patch.Symmetric_ind, seis_patch.Row_changed_by, seis_patch.Row_changed_date, seis_patch.Row_created_by, seis_patch.Row_effective_date, seis_patch.Row_expiry_date, seis_patch.Row_quality, seis_patch.Patch_id)
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

func PatchSeisPatch(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_patch set "
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
	query += " where patch_id = :id"

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

func DeleteSeisPatch(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_patch dto.Seis_patch
	seis_patch.Patch_id = id

	stmt, err := db.Prepare("delete from seis_patch where patch_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_patch.Patch_id)
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


