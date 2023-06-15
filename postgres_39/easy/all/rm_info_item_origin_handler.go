package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmInfoItemOrigin(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_info_item_origin")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_info_item_origin

	for rows.Next() {
		var rm_info_item_origin dto.Rm_info_item_origin
		if err := rows.Scan(&rm_info_item_origin.Info_item_subtype, &rm_info_item_origin.Information_item_id, &rm_info_item_origin.Origin_seq_no, &rm_info_item_origin.Active_ind, &rm_info_item_origin.Effective_date, &rm_info_item_origin.Expiry_date, &rm_info_item_origin.Information_process, &rm_info_item_origin.Parent_information_item_id, &rm_info_item_origin.Parent_info_item_subtype, &rm_info_item_origin.Ppdm_guid, &rm_info_item_origin.Processor, &rm_info_item_origin.Process_date, &rm_info_item_origin.Remark, &rm_info_item_origin.Source, &rm_info_item_origin.Version, &rm_info_item_origin.Row_changed_by, &rm_info_item_origin.Row_changed_date, &rm_info_item_origin.Row_created_by, &rm_info_item_origin.Row_created_date, &rm_info_item_origin.Row_effective_date, &rm_info_item_origin.Row_expiry_date, &rm_info_item_origin.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_info_item_origin to result
		result = append(result, rm_info_item_origin)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmInfoItemOrigin(c *fiber.Ctx) error {
	var rm_info_item_origin dto.Rm_info_item_origin

	setDefaults(&rm_info_item_origin)

	if err := json.Unmarshal(c.Body(), &rm_info_item_origin); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_info_item_origin values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	rm_info_item_origin.Row_created_date = formatDate(rm_info_item_origin.Row_created_date)
	rm_info_item_origin.Row_changed_date = formatDate(rm_info_item_origin.Row_changed_date)
	rm_info_item_origin.Effective_date = formatDateString(rm_info_item_origin.Effective_date)
	rm_info_item_origin.Expiry_date = formatDateString(rm_info_item_origin.Expiry_date)
	rm_info_item_origin.Process_date = formatDateString(rm_info_item_origin.Process_date)
	rm_info_item_origin.Row_effective_date = formatDateString(rm_info_item_origin.Row_effective_date)
	rm_info_item_origin.Row_expiry_date = formatDateString(rm_info_item_origin.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_origin.Info_item_subtype, rm_info_item_origin.Information_item_id, rm_info_item_origin.Origin_seq_no, rm_info_item_origin.Active_ind, rm_info_item_origin.Effective_date, rm_info_item_origin.Expiry_date, rm_info_item_origin.Information_process, rm_info_item_origin.Parent_information_item_id, rm_info_item_origin.Parent_info_item_subtype, rm_info_item_origin.Ppdm_guid, rm_info_item_origin.Processor, rm_info_item_origin.Process_date, rm_info_item_origin.Remark, rm_info_item_origin.Source, rm_info_item_origin.Version, rm_info_item_origin.Row_changed_by, rm_info_item_origin.Row_changed_date, rm_info_item_origin.Row_created_by, rm_info_item_origin.Row_created_date, rm_info_item_origin.Row_effective_date, rm_info_item_origin.Row_expiry_date, rm_info_item_origin.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmInfoItemOrigin(c *fiber.Ctx) error {
	var rm_info_item_origin dto.Rm_info_item_origin

	setDefaults(&rm_info_item_origin)

	if err := json.Unmarshal(c.Body(), &rm_info_item_origin); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_info_item_origin.Info_item_subtype = id

    if rm_info_item_origin.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_info_item_origin set information_item_id = :1, origin_seq_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, information_process = :6, parent_information_item_id = :7, parent_info_item_subtype = :8, ppdm_guid = :9, processor = :10, process_date = :11, remark = :12, source = :13, version = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where info_item_subtype = :22")
	if err != nil {
		return err
	}

	rm_info_item_origin.Row_changed_date = formatDate(rm_info_item_origin.Row_changed_date)
	rm_info_item_origin.Effective_date = formatDateString(rm_info_item_origin.Effective_date)
	rm_info_item_origin.Expiry_date = formatDateString(rm_info_item_origin.Expiry_date)
	rm_info_item_origin.Process_date = formatDateString(rm_info_item_origin.Process_date)
	rm_info_item_origin.Row_effective_date = formatDateString(rm_info_item_origin.Row_effective_date)
	rm_info_item_origin.Row_expiry_date = formatDateString(rm_info_item_origin.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_origin.Information_item_id, rm_info_item_origin.Origin_seq_no, rm_info_item_origin.Active_ind, rm_info_item_origin.Effective_date, rm_info_item_origin.Expiry_date, rm_info_item_origin.Information_process, rm_info_item_origin.Parent_information_item_id, rm_info_item_origin.Parent_info_item_subtype, rm_info_item_origin.Ppdm_guid, rm_info_item_origin.Processor, rm_info_item_origin.Process_date, rm_info_item_origin.Remark, rm_info_item_origin.Source, rm_info_item_origin.Version, rm_info_item_origin.Row_changed_by, rm_info_item_origin.Row_changed_date, rm_info_item_origin.Row_created_by, rm_info_item_origin.Row_effective_date, rm_info_item_origin.Row_expiry_date, rm_info_item_origin.Row_quality, rm_info_item_origin.Info_item_subtype)
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

func PatchRmInfoItemOrigin(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_info_item_origin set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "process_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where info_item_subtype = :id"

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

func DeleteRmInfoItemOrigin(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_info_item_origin dto.Rm_info_item_origin
	rm_info_item_origin.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_info_item_origin where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_info_item_origin.Info_item_subtype)
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


