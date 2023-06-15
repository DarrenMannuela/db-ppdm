package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmPhysItemStore(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_phys_item_store")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_phys_item_store

	for rows.Next() {
		var rm_phys_item_store dto.Rm_phys_item_store
		if err := rows.Scan(&rm_phys_item_store.Store_id, &rm_phys_item_store.Physical_item_id, &rm_phys_item_store.Store_seq_no, &rm_phys_item_store.Active_ind, &rm_phys_item_store.Create_date, &rm_phys_item_store.Effective_date, &rm_phys_item_store.Expiry_date, &rm_phys_item_store.Physical_item_status, &rm_phys_item_store.Ppdm_guid, &rm_phys_item_store.Preferred_location_ind, &rm_phys_item_store.Remark, &rm_phys_item_store.Source, &rm_phys_item_store.Row_changed_by, &rm_phys_item_store.Row_changed_date, &rm_phys_item_store.Row_created_by, &rm_phys_item_store.Row_created_date, &rm_phys_item_store.Row_effective_date, &rm_phys_item_store.Row_expiry_date, &rm_phys_item_store.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_phys_item_store to result
		result = append(result, rm_phys_item_store)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmPhysItemStore(c *fiber.Ctx) error {
	var rm_phys_item_store dto.Rm_phys_item_store

	setDefaults(&rm_phys_item_store)

	if err := json.Unmarshal(c.Body(), &rm_phys_item_store); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_phys_item_store values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	rm_phys_item_store.Row_created_date = formatDate(rm_phys_item_store.Row_created_date)
	rm_phys_item_store.Row_changed_date = formatDate(rm_phys_item_store.Row_changed_date)
	rm_phys_item_store.Create_date = formatDateString(rm_phys_item_store.Create_date)
	rm_phys_item_store.Effective_date = formatDateString(rm_phys_item_store.Effective_date)
	rm_phys_item_store.Expiry_date = formatDateString(rm_phys_item_store.Expiry_date)
	rm_phys_item_store.Row_effective_date = formatDateString(rm_phys_item_store.Row_effective_date)
	rm_phys_item_store.Row_expiry_date = formatDateString(rm_phys_item_store.Row_expiry_date)






	rows, err := stmt.Exec(rm_phys_item_store.Store_id, rm_phys_item_store.Physical_item_id, rm_phys_item_store.Store_seq_no, rm_phys_item_store.Active_ind, rm_phys_item_store.Create_date, rm_phys_item_store.Effective_date, rm_phys_item_store.Expiry_date, rm_phys_item_store.Physical_item_status, rm_phys_item_store.Ppdm_guid, rm_phys_item_store.Preferred_location_ind, rm_phys_item_store.Remark, rm_phys_item_store.Source, rm_phys_item_store.Row_changed_by, rm_phys_item_store.Row_changed_date, rm_phys_item_store.Row_created_by, rm_phys_item_store.Row_created_date, rm_phys_item_store.Row_effective_date, rm_phys_item_store.Row_expiry_date, rm_phys_item_store.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmPhysItemStore(c *fiber.Ctx) error {
	var rm_phys_item_store dto.Rm_phys_item_store

	setDefaults(&rm_phys_item_store)

	if err := json.Unmarshal(c.Body(), &rm_phys_item_store); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_phys_item_store.Store_id = id

    if rm_phys_item_store.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_phys_item_store set physical_item_id = :1, store_seq_no = :2, active_ind = :3, create_date = :4, effective_date = :5, expiry_date = :6, physical_item_status = :7, ppdm_guid = :8, preferred_location_ind = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where store_id = :19")
	if err != nil {
		return err
	}

	rm_phys_item_store.Row_changed_date = formatDate(rm_phys_item_store.Row_changed_date)
	rm_phys_item_store.Create_date = formatDateString(rm_phys_item_store.Create_date)
	rm_phys_item_store.Effective_date = formatDateString(rm_phys_item_store.Effective_date)
	rm_phys_item_store.Expiry_date = formatDateString(rm_phys_item_store.Expiry_date)
	rm_phys_item_store.Row_effective_date = formatDateString(rm_phys_item_store.Row_effective_date)
	rm_phys_item_store.Row_expiry_date = formatDateString(rm_phys_item_store.Row_expiry_date)






	rows, err := stmt.Exec(rm_phys_item_store.Physical_item_id, rm_phys_item_store.Store_seq_no, rm_phys_item_store.Active_ind, rm_phys_item_store.Create_date, rm_phys_item_store.Effective_date, rm_phys_item_store.Expiry_date, rm_phys_item_store.Physical_item_status, rm_phys_item_store.Ppdm_guid, rm_phys_item_store.Preferred_location_ind, rm_phys_item_store.Remark, rm_phys_item_store.Source, rm_phys_item_store.Row_changed_by, rm_phys_item_store.Row_changed_date, rm_phys_item_store.Row_created_by, rm_phys_item_store.Row_effective_date, rm_phys_item_store.Row_expiry_date, rm_phys_item_store.Row_quality, rm_phys_item_store.Store_id)
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

func PatchRmPhysItemStore(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_phys_item_store set "
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
		} else if key == "create_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteRmPhysItemStore(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_phys_item_store dto.Rm_phys_item_store
	rm_phys_item_store.Store_id = id

	stmt, err := db.Prepare("delete from rm_phys_item_store where store_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_phys_item_store.Store_id)
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


