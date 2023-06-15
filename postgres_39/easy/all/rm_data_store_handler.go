package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmDataStore(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_data_store")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_data_store

	for rows.Next() {
		var rm_data_store dto.Rm_data_store
		if err := rows.Scan(&rm_data_store.Store_id, &rm_data_store.Active_ind, &rm_data_store.Address_obs_no, &rm_data_store.Address_source, &rm_data_store.Business_associate_id, &rm_data_store.Contained_by_store_id, &rm_data_store.Data_store_hier_id, &rm_data_store.Data_store_name, &rm_data_store.Data_store_type, &rm_data_store.Effective_date, &rm_data_store.Expiry_date, &rm_data_store.Hier_level_id, &rm_data_store.Location_id, &rm_data_store.Long_location, &rm_data_store.Physical_store_status, &rm_data_store.Ppdm_guid, &rm_data_store.Remark, &rm_data_store.Short_location, &rm_data_store.Source, &rm_data_store.Total_capacity, &rm_data_store.Used_capacity, &rm_data_store.Used_capacity_date, &rm_data_store.Row_changed_by, &rm_data_store.Row_changed_date, &rm_data_store.Row_created_by, &rm_data_store.Row_created_date, &rm_data_store.Row_effective_date, &rm_data_store.Row_expiry_date, &rm_data_store.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_data_store to result
		result = append(result, rm_data_store)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmDataStore(c *fiber.Ctx) error {
	var rm_data_store dto.Rm_data_store

	setDefaults(&rm_data_store)

	if err := json.Unmarshal(c.Body(), &rm_data_store); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_data_store values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	rm_data_store.Row_created_date = formatDate(rm_data_store.Row_created_date)
	rm_data_store.Row_changed_date = formatDate(rm_data_store.Row_changed_date)
	rm_data_store.Effective_date = formatDateString(rm_data_store.Effective_date)
	rm_data_store.Expiry_date = formatDateString(rm_data_store.Expiry_date)
	rm_data_store.Used_capacity_date = formatDateString(rm_data_store.Used_capacity_date)
	rm_data_store.Row_effective_date = formatDateString(rm_data_store.Row_effective_date)
	rm_data_store.Row_expiry_date = formatDateString(rm_data_store.Row_expiry_date)






	rows, err := stmt.Exec(rm_data_store.Store_id, rm_data_store.Active_ind, rm_data_store.Address_obs_no, rm_data_store.Address_source, rm_data_store.Business_associate_id, rm_data_store.Contained_by_store_id, rm_data_store.Data_store_hier_id, rm_data_store.Data_store_name, rm_data_store.Data_store_type, rm_data_store.Effective_date, rm_data_store.Expiry_date, rm_data_store.Hier_level_id, rm_data_store.Location_id, rm_data_store.Long_location, rm_data_store.Physical_store_status, rm_data_store.Ppdm_guid, rm_data_store.Remark, rm_data_store.Short_location, rm_data_store.Source, rm_data_store.Total_capacity, rm_data_store.Used_capacity, rm_data_store.Used_capacity_date, rm_data_store.Row_changed_by, rm_data_store.Row_changed_date, rm_data_store.Row_created_by, rm_data_store.Row_created_date, rm_data_store.Row_effective_date, rm_data_store.Row_expiry_date, rm_data_store.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmDataStore(c *fiber.Ctx) error {
	var rm_data_store dto.Rm_data_store

	setDefaults(&rm_data_store)

	if err := json.Unmarshal(c.Body(), &rm_data_store); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_data_store.Store_id = id

    if rm_data_store.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_data_store set active_ind = :1, address_obs_no = :2, address_source = :3, business_associate_id = :4, contained_by_store_id = :5, data_store_hier_id = :6, data_store_name = :7, data_store_type = :8, effective_date = :9, expiry_date = :10, hier_level_id = :11, location_id = :12, long_location = :13, physical_store_status = :14, ppdm_guid = :15, remark = :16, short_location = :17, source = :18, total_capacity = :19, used_capacity = :20, used_capacity_date = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where store_id = :29")
	if err != nil {
		return err
	}

	rm_data_store.Row_changed_date = formatDate(rm_data_store.Row_changed_date)
	rm_data_store.Effective_date = formatDateString(rm_data_store.Effective_date)
	rm_data_store.Expiry_date = formatDateString(rm_data_store.Expiry_date)
	rm_data_store.Used_capacity_date = formatDateString(rm_data_store.Used_capacity_date)
	rm_data_store.Row_effective_date = formatDateString(rm_data_store.Row_effective_date)
	rm_data_store.Row_expiry_date = formatDateString(rm_data_store.Row_expiry_date)






	rows, err := stmt.Exec(rm_data_store.Active_ind, rm_data_store.Address_obs_no, rm_data_store.Address_source, rm_data_store.Business_associate_id, rm_data_store.Contained_by_store_id, rm_data_store.Data_store_hier_id, rm_data_store.Data_store_name, rm_data_store.Data_store_type, rm_data_store.Effective_date, rm_data_store.Expiry_date, rm_data_store.Hier_level_id, rm_data_store.Location_id, rm_data_store.Long_location, rm_data_store.Physical_store_status, rm_data_store.Ppdm_guid, rm_data_store.Remark, rm_data_store.Short_location, rm_data_store.Source, rm_data_store.Total_capacity, rm_data_store.Used_capacity, rm_data_store.Used_capacity_date, rm_data_store.Row_changed_by, rm_data_store.Row_changed_date, rm_data_store.Row_created_by, rm_data_store.Row_effective_date, rm_data_store.Row_expiry_date, rm_data_store.Row_quality, rm_data_store.Store_id)
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

func PatchRmDataStore(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_data_store set "
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

func DeleteRmDataStore(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_data_store dto.Rm_data_store
	rm_data_store.Store_id = id

	stmt, err := db.Prepare("delete from rm_data_store where store_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_data_store.Store_id)
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


