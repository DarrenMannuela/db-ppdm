package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetReserveEntity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from reserve_entity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Reserve_entity

	for rows.Next() {
		var reserve_entity dto.Reserve_entity
		if err := rows.Scan(&reserve_entity.Resent_id, &reserve_entity.Active_ind, &reserve_entity.Created_by_ba_id, &reserve_entity.Description, &reserve_entity.Effective_date, &reserve_entity.Expiry_date, &reserve_entity.Group_type, &reserve_entity.Last_approve_ba_id, &reserve_entity.Last_update_ba_id, &reserve_entity.Last_update_date, &reserve_entity.Ppdm_guid, &reserve_entity.Primary_product_type, &reserve_entity.Remark, &reserve_entity.Report_ind, &reserve_entity.Resent_long_name, &reserve_entity.Resent_short_name, &reserve_entity.Source, &reserve_entity.Update_schedule, &reserve_entity.Row_changed_by, &reserve_entity.Row_changed_date, &reserve_entity.Row_created_by, &reserve_entity.Row_created_date, &reserve_entity.Row_effective_date, &reserve_entity.Row_expiry_date, &reserve_entity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append reserve_entity to result
		result = append(result, reserve_entity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetReserveEntity(c *fiber.Ctx) error {
	var reserve_entity dto.Reserve_entity

	setDefaults(&reserve_entity)

	if err := json.Unmarshal(c.Body(), &reserve_entity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into reserve_entity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	reserve_entity.Row_created_date = formatDate(reserve_entity.Row_created_date)
	reserve_entity.Row_changed_date = formatDate(reserve_entity.Row_changed_date)
	reserve_entity.Effective_date = formatDateString(reserve_entity.Effective_date)
	reserve_entity.Expiry_date = formatDateString(reserve_entity.Expiry_date)
	reserve_entity.Last_update_date = formatDateString(reserve_entity.Last_update_date)
	reserve_entity.Row_effective_date = formatDateString(reserve_entity.Row_effective_date)
	reserve_entity.Row_expiry_date = formatDateString(reserve_entity.Row_expiry_date)






	rows, err := stmt.Exec(reserve_entity.Resent_id, reserve_entity.Active_ind, reserve_entity.Created_by_ba_id, reserve_entity.Description, reserve_entity.Effective_date, reserve_entity.Expiry_date, reserve_entity.Group_type, reserve_entity.Last_approve_ba_id, reserve_entity.Last_update_ba_id, reserve_entity.Last_update_date, reserve_entity.Ppdm_guid, reserve_entity.Primary_product_type, reserve_entity.Remark, reserve_entity.Report_ind, reserve_entity.Resent_long_name, reserve_entity.Resent_short_name, reserve_entity.Source, reserve_entity.Update_schedule, reserve_entity.Row_changed_by, reserve_entity.Row_changed_date, reserve_entity.Row_created_by, reserve_entity.Row_created_date, reserve_entity.Row_effective_date, reserve_entity.Row_expiry_date, reserve_entity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateReserveEntity(c *fiber.Ctx) error {
	var reserve_entity dto.Reserve_entity

	setDefaults(&reserve_entity)

	if err := json.Unmarshal(c.Body(), &reserve_entity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	reserve_entity.Resent_id = id

    if reserve_entity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update reserve_entity set active_ind = :1, created_by_ba_id = :2, description = :3, effective_date = :4, expiry_date = :5, group_type = :6, last_approve_ba_id = :7, last_update_ba_id = :8, last_update_date = :9, ppdm_guid = :10, primary_product_type = :11, remark = :12, report_ind = :13, resent_long_name = :14, resent_short_name = :15, source = :16, update_schedule = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where resent_id = :25")
	if err != nil {
		return err
	}

	reserve_entity.Row_changed_date = formatDate(reserve_entity.Row_changed_date)
	reserve_entity.Effective_date = formatDateString(reserve_entity.Effective_date)
	reserve_entity.Expiry_date = formatDateString(reserve_entity.Expiry_date)
	reserve_entity.Last_update_date = formatDateString(reserve_entity.Last_update_date)
	reserve_entity.Row_effective_date = formatDateString(reserve_entity.Row_effective_date)
	reserve_entity.Row_expiry_date = formatDateString(reserve_entity.Row_expiry_date)






	rows, err := stmt.Exec(reserve_entity.Active_ind, reserve_entity.Created_by_ba_id, reserve_entity.Description, reserve_entity.Effective_date, reserve_entity.Expiry_date, reserve_entity.Group_type, reserve_entity.Last_approve_ba_id, reserve_entity.Last_update_ba_id, reserve_entity.Last_update_date, reserve_entity.Ppdm_guid, reserve_entity.Primary_product_type, reserve_entity.Remark, reserve_entity.Report_ind, reserve_entity.Resent_long_name, reserve_entity.Resent_short_name, reserve_entity.Source, reserve_entity.Update_schedule, reserve_entity.Row_changed_by, reserve_entity.Row_changed_date, reserve_entity.Row_created_by, reserve_entity.Row_effective_date, reserve_entity.Row_expiry_date, reserve_entity.Row_quality, reserve_entity.Resent_id)
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

func PatchReserveEntity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update reserve_entity set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "last_update_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where resent_id = :id"

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

func DeleteReserveEntity(c *fiber.Ctx) error {
	id := c.Params("id")
	var reserve_entity dto.Reserve_entity
	reserve_entity.Resent_id = id

	stmt, err := db.Prepare("delete from reserve_entity where resent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(reserve_entity.Resent_id)
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


