package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetReserveClass(c *fiber.Ctx) error {
	rows, err := db.Query("select * from reserve_class")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Reserve_class

	for rows.Next() {
		var reserve_class dto.Reserve_class
		if err := rows.Scan(&reserve_class.Reserve_class_id, &reserve_class.Abbreviation, &reserve_class.Active_ind, &reserve_class.Confidence_type, &reserve_class.Developed_ind, &reserve_class.Effective_date, &reserve_class.Expiry_date, &reserve_class.Long_name, &reserve_class.Owner_ba_id, &reserve_class.Ppdm_guid, &reserve_class.Preferred_ind, &reserve_class.Production_ind, &reserve_class.Remark, &reserve_class.Short_name, &reserve_class.Source, &reserve_class.Use_type, &reserve_class.Row_changed_by, &reserve_class.Row_changed_date, &reserve_class.Row_created_by, &reserve_class.Row_created_date, &reserve_class.Row_effective_date, &reserve_class.Row_expiry_date, &reserve_class.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append reserve_class to result
		result = append(result, reserve_class)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetReserveClass(c *fiber.Ctx) error {
	var reserve_class dto.Reserve_class

	setDefaults(&reserve_class)

	if err := json.Unmarshal(c.Body(), &reserve_class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into reserve_class values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	reserve_class.Row_created_date = formatDate(reserve_class.Row_created_date)
	reserve_class.Row_changed_date = formatDate(reserve_class.Row_changed_date)
	reserve_class.Effective_date = formatDateString(reserve_class.Effective_date)
	reserve_class.Expiry_date = formatDateString(reserve_class.Expiry_date)
	reserve_class.Row_effective_date = formatDateString(reserve_class.Row_effective_date)
	reserve_class.Row_expiry_date = formatDateString(reserve_class.Row_expiry_date)






	rows, err := stmt.Exec(reserve_class.Reserve_class_id, reserve_class.Abbreviation, reserve_class.Active_ind, reserve_class.Confidence_type, reserve_class.Developed_ind, reserve_class.Effective_date, reserve_class.Expiry_date, reserve_class.Long_name, reserve_class.Owner_ba_id, reserve_class.Ppdm_guid, reserve_class.Preferred_ind, reserve_class.Production_ind, reserve_class.Remark, reserve_class.Short_name, reserve_class.Source, reserve_class.Use_type, reserve_class.Row_changed_by, reserve_class.Row_changed_date, reserve_class.Row_created_by, reserve_class.Row_created_date, reserve_class.Row_effective_date, reserve_class.Row_expiry_date, reserve_class.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateReserveClass(c *fiber.Ctx) error {
	var reserve_class dto.Reserve_class

	setDefaults(&reserve_class)

	if err := json.Unmarshal(c.Body(), &reserve_class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	reserve_class.Reserve_class_id = id

    if reserve_class.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update reserve_class set abbreviation = :1, active_ind = :2, confidence_type = :3, developed_ind = :4, effective_date = :5, expiry_date = :6, long_name = :7, owner_ba_id = :8, ppdm_guid = :9, preferred_ind = :10, production_ind = :11, remark = :12, short_name = :13, source = :14, use_type = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where reserve_class_id = :23")
	if err != nil {
		return err
	}

	reserve_class.Row_changed_date = formatDate(reserve_class.Row_changed_date)
	reserve_class.Effective_date = formatDateString(reserve_class.Effective_date)
	reserve_class.Expiry_date = formatDateString(reserve_class.Expiry_date)
	reserve_class.Row_effective_date = formatDateString(reserve_class.Row_effective_date)
	reserve_class.Row_expiry_date = formatDateString(reserve_class.Row_expiry_date)






	rows, err := stmt.Exec(reserve_class.Abbreviation, reserve_class.Active_ind, reserve_class.Confidence_type, reserve_class.Developed_ind, reserve_class.Effective_date, reserve_class.Expiry_date, reserve_class.Long_name, reserve_class.Owner_ba_id, reserve_class.Ppdm_guid, reserve_class.Preferred_ind, reserve_class.Production_ind, reserve_class.Remark, reserve_class.Short_name, reserve_class.Source, reserve_class.Use_type, reserve_class.Row_changed_by, reserve_class.Row_changed_date, reserve_class.Row_created_by, reserve_class.Row_effective_date, reserve_class.Row_expiry_date, reserve_class.Row_quality, reserve_class.Reserve_class_id)
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

func PatchReserveClass(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update reserve_class set "
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
	query += " where reserve_class_id = :id"

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

func DeleteReserveClass(c *fiber.Ctx) error {
	id := c.Params("id")
	var reserve_class dto.Reserve_class
	reserve_class.Reserve_class_id = id

	stmt, err := db.Prepare("delete from reserve_class where reserve_class_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(reserve_class.Reserve_class_id)
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


