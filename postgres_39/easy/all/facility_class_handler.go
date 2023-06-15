package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityClass(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_class")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_class

	for rows.Next() {
		var facility_class dto.Facility_class
		if err := rows.Scan(&facility_class.Facility_id, &facility_class.Facility_type, &facility_class.Facility_class_type, &facility_class.Facility_class_seq_no, &facility_class.Active_ind, &facility_class.Effective_date, &facility_class.Expiry_date, &facility_class.Ppdm_guid, &facility_class.Remark, &facility_class.Source, &facility_class.Row_changed_by, &facility_class.Row_changed_date, &facility_class.Row_created_by, &facility_class.Row_created_date, &facility_class.Row_effective_date, &facility_class.Row_expiry_date, &facility_class.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_class to result
		result = append(result, facility_class)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityClass(c *fiber.Ctx) error {
	var facility_class dto.Facility_class

	setDefaults(&facility_class)

	if err := json.Unmarshal(c.Body(), &facility_class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_class values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	facility_class.Row_created_date = formatDate(facility_class.Row_created_date)
	facility_class.Row_changed_date = formatDate(facility_class.Row_changed_date)
	facility_class.Effective_date = formatDateString(facility_class.Effective_date)
	facility_class.Expiry_date = formatDateString(facility_class.Expiry_date)
	facility_class.Row_effective_date = formatDateString(facility_class.Row_effective_date)
	facility_class.Row_expiry_date = formatDateString(facility_class.Row_expiry_date)






	rows, err := stmt.Exec(facility_class.Facility_id, facility_class.Facility_type, facility_class.Facility_class_type, facility_class.Facility_class_seq_no, facility_class.Active_ind, facility_class.Effective_date, facility_class.Expiry_date, facility_class.Ppdm_guid, facility_class.Remark, facility_class.Source, facility_class.Row_changed_by, facility_class.Row_changed_date, facility_class.Row_created_by, facility_class.Row_created_date, facility_class.Row_effective_date, facility_class.Row_expiry_date, facility_class.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityClass(c *fiber.Ctx) error {
	var facility_class dto.Facility_class

	setDefaults(&facility_class)

	if err := json.Unmarshal(c.Body(), &facility_class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_class.Facility_id = id

    if facility_class.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_class set facility_type = :1, facility_class_type = :2, facility_class_seq_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where facility_id = :17")
	if err != nil {
		return err
	}

	facility_class.Row_changed_date = formatDate(facility_class.Row_changed_date)
	facility_class.Effective_date = formatDateString(facility_class.Effective_date)
	facility_class.Expiry_date = formatDateString(facility_class.Expiry_date)
	facility_class.Row_effective_date = formatDateString(facility_class.Row_effective_date)
	facility_class.Row_expiry_date = formatDateString(facility_class.Row_expiry_date)






	rows, err := stmt.Exec(facility_class.Facility_type, facility_class.Facility_class_type, facility_class.Facility_class_seq_no, facility_class.Active_ind, facility_class.Effective_date, facility_class.Expiry_date, facility_class.Ppdm_guid, facility_class.Remark, facility_class.Source, facility_class.Row_changed_by, facility_class.Row_changed_date, facility_class.Row_created_by, facility_class.Row_effective_date, facility_class.Row_expiry_date, facility_class.Row_quality, facility_class.Facility_id)
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

func PatchFacilityClass(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_class set "
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
	query += " where facility_id = :id"

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

func DeleteFacilityClass(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_class dto.Facility_class
	facility_class.Facility_id = id

	stmt, err := db.Prepare("delete from facility_class where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_class.Facility_id)
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


