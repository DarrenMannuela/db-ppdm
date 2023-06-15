package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmDataStoreMedia(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_data_store_media")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_data_store_media

	for rows.Next() {
		var rm_data_store_media dto.Rm_data_store_media
		if err := rows.Scan(&rm_data_store_media.Store_id, &rm_data_store_media.Structure_obs_no, &rm_data_store_media.Media_obs_no, &rm_data_store_media.Active_ind, &rm_data_store_media.Effective_date, &rm_data_store_media.Expiry_date, &rm_data_store_media.Media_type, &rm_data_store_media.Ppdm_guid, &rm_data_store_media.Remark, &rm_data_store_media.Source, &rm_data_store_media.Total_capacity, &rm_data_store_media.Used_capacity, &rm_data_store_media.Used_capacity_date, &rm_data_store_media.Row_changed_by, &rm_data_store_media.Row_changed_date, &rm_data_store_media.Row_created_by, &rm_data_store_media.Row_created_date, &rm_data_store_media.Row_effective_date, &rm_data_store_media.Row_expiry_date, &rm_data_store_media.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_data_store_media to result
		result = append(result, rm_data_store_media)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmDataStoreMedia(c *fiber.Ctx) error {
	var rm_data_store_media dto.Rm_data_store_media

	setDefaults(&rm_data_store_media)

	if err := json.Unmarshal(c.Body(), &rm_data_store_media); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_data_store_media values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	rm_data_store_media.Row_created_date = formatDate(rm_data_store_media.Row_created_date)
	rm_data_store_media.Row_changed_date = formatDate(rm_data_store_media.Row_changed_date)
	rm_data_store_media.Effective_date = formatDateString(rm_data_store_media.Effective_date)
	rm_data_store_media.Expiry_date = formatDateString(rm_data_store_media.Expiry_date)
	rm_data_store_media.Used_capacity_date = formatDateString(rm_data_store_media.Used_capacity_date)
	rm_data_store_media.Row_effective_date = formatDateString(rm_data_store_media.Row_effective_date)
	rm_data_store_media.Row_expiry_date = formatDateString(rm_data_store_media.Row_expiry_date)






	rows, err := stmt.Exec(rm_data_store_media.Store_id, rm_data_store_media.Structure_obs_no, rm_data_store_media.Media_obs_no, rm_data_store_media.Active_ind, rm_data_store_media.Effective_date, rm_data_store_media.Expiry_date, rm_data_store_media.Media_type, rm_data_store_media.Ppdm_guid, rm_data_store_media.Remark, rm_data_store_media.Source, rm_data_store_media.Total_capacity, rm_data_store_media.Used_capacity, rm_data_store_media.Used_capacity_date, rm_data_store_media.Row_changed_by, rm_data_store_media.Row_changed_date, rm_data_store_media.Row_created_by, rm_data_store_media.Row_created_date, rm_data_store_media.Row_effective_date, rm_data_store_media.Row_expiry_date, rm_data_store_media.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmDataStoreMedia(c *fiber.Ctx) error {
	var rm_data_store_media dto.Rm_data_store_media

	setDefaults(&rm_data_store_media)

	if err := json.Unmarshal(c.Body(), &rm_data_store_media); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_data_store_media.Store_id = id

    if rm_data_store_media.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_data_store_media set structure_obs_no = :1, media_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, media_type = :6, ppdm_guid = :7, remark = :8, source = :9, total_capacity = :10, used_capacity = :11, used_capacity_date = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where store_id = :20")
	if err != nil {
		return err
	}

	rm_data_store_media.Row_changed_date = formatDate(rm_data_store_media.Row_changed_date)
	rm_data_store_media.Effective_date = formatDateString(rm_data_store_media.Effective_date)
	rm_data_store_media.Expiry_date = formatDateString(rm_data_store_media.Expiry_date)
	rm_data_store_media.Used_capacity_date = formatDateString(rm_data_store_media.Used_capacity_date)
	rm_data_store_media.Row_effective_date = formatDateString(rm_data_store_media.Row_effective_date)
	rm_data_store_media.Row_expiry_date = formatDateString(rm_data_store_media.Row_expiry_date)






	rows, err := stmt.Exec(rm_data_store_media.Structure_obs_no, rm_data_store_media.Media_obs_no, rm_data_store_media.Active_ind, rm_data_store_media.Effective_date, rm_data_store_media.Expiry_date, rm_data_store_media.Media_type, rm_data_store_media.Ppdm_guid, rm_data_store_media.Remark, rm_data_store_media.Source, rm_data_store_media.Total_capacity, rm_data_store_media.Used_capacity, rm_data_store_media.Used_capacity_date, rm_data_store_media.Row_changed_by, rm_data_store_media.Row_changed_date, rm_data_store_media.Row_created_by, rm_data_store_media.Row_effective_date, rm_data_store_media.Row_expiry_date, rm_data_store_media.Row_quality, rm_data_store_media.Store_id)
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

func PatchRmDataStoreMedia(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_data_store_media set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "used_capacity_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where store_id = :id"

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

func DeleteRmDataStoreMedia(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_data_store_media dto.Rm_data_store_media
	rm_data_store_media.Store_id = id

	stmt, err := db.Prepare("delete from rm_data_store_media where store_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_data_store_media.Store_id)
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


