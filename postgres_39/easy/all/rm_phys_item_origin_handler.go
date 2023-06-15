package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmPhysItemOrigin(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_phys_item_origin")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_phys_item_origin

	for rows.Next() {
		var rm_phys_item_origin dto.Rm_phys_item_origin
		if err := rows.Scan(&rm_phys_item_origin.Physical_item_id, &rm_phys_item_origin.Origin_seq_no, &rm_phys_item_origin.Active_ind, &rm_phys_item_origin.Business_associate_id, &rm_phys_item_origin.Effective_date, &rm_phys_item_origin.Expiry_date, &rm_phys_item_origin.Parent_physical_item_id, &rm_phys_item_origin.Physical_process, &rm_phys_item_origin.Ppdm_guid, &rm_phys_item_origin.Process_date, &rm_phys_item_origin.Remark, &rm_phys_item_origin.Source, &rm_phys_item_origin.Version, &rm_phys_item_origin.Row_changed_by, &rm_phys_item_origin.Row_changed_date, &rm_phys_item_origin.Row_created_by, &rm_phys_item_origin.Row_created_date, &rm_phys_item_origin.Row_effective_date, &rm_phys_item_origin.Row_expiry_date, &rm_phys_item_origin.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_phys_item_origin to result
		result = append(result, rm_phys_item_origin)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmPhysItemOrigin(c *fiber.Ctx) error {
	var rm_phys_item_origin dto.Rm_phys_item_origin

	setDefaults(&rm_phys_item_origin)

	if err := json.Unmarshal(c.Body(), &rm_phys_item_origin); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_phys_item_origin values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	rm_phys_item_origin.Row_created_date = formatDate(rm_phys_item_origin.Row_created_date)
	rm_phys_item_origin.Row_changed_date = formatDate(rm_phys_item_origin.Row_changed_date)
	rm_phys_item_origin.Effective_date = formatDateString(rm_phys_item_origin.Effective_date)
	rm_phys_item_origin.Expiry_date = formatDateString(rm_phys_item_origin.Expiry_date)
	rm_phys_item_origin.Process_date = formatDateString(rm_phys_item_origin.Process_date)
	rm_phys_item_origin.Row_effective_date = formatDateString(rm_phys_item_origin.Row_effective_date)
	rm_phys_item_origin.Row_expiry_date = formatDateString(rm_phys_item_origin.Row_expiry_date)






	rows, err := stmt.Exec(rm_phys_item_origin.Physical_item_id, rm_phys_item_origin.Origin_seq_no, rm_phys_item_origin.Active_ind, rm_phys_item_origin.Business_associate_id, rm_phys_item_origin.Effective_date, rm_phys_item_origin.Expiry_date, rm_phys_item_origin.Parent_physical_item_id, rm_phys_item_origin.Physical_process, rm_phys_item_origin.Ppdm_guid, rm_phys_item_origin.Process_date, rm_phys_item_origin.Remark, rm_phys_item_origin.Source, rm_phys_item_origin.Version, rm_phys_item_origin.Row_changed_by, rm_phys_item_origin.Row_changed_date, rm_phys_item_origin.Row_created_by, rm_phys_item_origin.Row_created_date, rm_phys_item_origin.Row_effective_date, rm_phys_item_origin.Row_expiry_date, rm_phys_item_origin.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmPhysItemOrigin(c *fiber.Ctx) error {
	var rm_phys_item_origin dto.Rm_phys_item_origin

	setDefaults(&rm_phys_item_origin)

	if err := json.Unmarshal(c.Body(), &rm_phys_item_origin); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_phys_item_origin.Physical_item_id = id

    if rm_phys_item_origin.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_phys_item_origin set origin_seq_no = :1, active_ind = :2, business_associate_id = :3, effective_date = :4, expiry_date = :5, parent_physical_item_id = :6, physical_process = :7, ppdm_guid = :8, process_date = :9, remark = :10, source = :11, version = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where physical_item_id = :20")
	if err != nil {
		return err
	}

	rm_phys_item_origin.Row_changed_date = formatDate(rm_phys_item_origin.Row_changed_date)
	rm_phys_item_origin.Effective_date = formatDateString(rm_phys_item_origin.Effective_date)
	rm_phys_item_origin.Expiry_date = formatDateString(rm_phys_item_origin.Expiry_date)
	rm_phys_item_origin.Process_date = formatDateString(rm_phys_item_origin.Process_date)
	rm_phys_item_origin.Row_effective_date = formatDateString(rm_phys_item_origin.Row_effective_date)
	rm_phys_item_origin.Row_expiry_date = formatDateString(rm_phys_item_origin.Row_expiry_date)






	rows, err := stmt.Exec(rm_phys_item_origin.Origin_seq_no, rm_phys_item_origin.Active_ind, rm_phys_item_origin.Business_associate_id, rm_phys_item_origin.Effective_date, rm_phys_item_origin.Expiry_date, rm_phys_item_origin.Parent_physical_item_id, rm_phys_item_origin.Physical_process, rm_phys_item_origin.Ppdm_guid, rm_phys_item_origin.Process_date, rm_phys_item_origin.Remark, rm_phys_item_origin.Source, rm_phys_item_origin.Version, rm_phys_item_origin.Row_changed_by, rm_phys_item_origin.Row_changed_date, rm_phys_item_origin.Row_created_by, rm_phys_item_origin.Row_effective_date, rm_phys_item_origin.Row_expiry_date, rm_phys_item_origin.Row_quality, rm_phys_item_origin.Physical_item_id)
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

func PatchRmPhysItemOrigin(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_phys_item_origin set "
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
	query += " where physical_item_id = :id"

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

func DeleteRmPhysItemOrigin(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_phys_item_origin dto.Rm_phys_item_origin
	rm_phys_item_origin.Physical_item_id = id

	stmt, err := db.Prepare("delete from rm_phys_item_origin where physical_item_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_phys_item_origin.Physical_item_id)
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


