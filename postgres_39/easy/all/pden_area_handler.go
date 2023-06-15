package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_area

	for rows.Next() {
		var pden_area dto.Pden_area
		if err := rows.Scan(&pden_area.Pden_subtype, &pden_area.Pden_id, &pden_area.Pden_source, &pden_area.Area_id, &pden_area.Area_type, &pden_area.Active_ind, &pden_area.Effective_date, &pden_area.Expiry_date, &pden_area.No_of_gas_wells, &pden_area.No_of_injection_wells, &pden_area.No_of_oil_wells, &pden_area.Ppdm_guid, &pden_area.Remark, &pden_area.Row_changed_by, &pden_area.Row_changed_date, &pden_area.Row_created_by, &pden_area.Row_created_date, &pden_area.Row_effective_date, &pden_area.Row_expiry_date, &pden_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_area to result
		result = append(result, pden_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenArea(c *fiber.Ctx) error {
	var pden_area dto.Pden_area

	setDefaults(&pden_area)

	if err := json.Unmarshal(c.Body(), &pden_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	pden_area.Row_created_date = formatDate(pden_area.Row_created_date)
	pden_area.Row_changed_date = formatDate(pden_area.Row_changed_date)
	pden_area.Effective_date = formatDateString(pden_area.Effective_date)
	pden_area.Expiry_date = formatDateString(pden_area.Expiry_date)
	pden_area.Row_effective_date = formatDateString(pden_area.Row_effective_date)
	pden_area.Row_expiry_date = formatDateString(pden_area.Row_expiry_date)






	rows, err := stmt.Exec(pden_area.Pden_subtype, pden_area.Pden_id, pden_area.Pden_source, pden_area.Area_id, pden_area.Area_type, pden_area.Active_ind, pden_area.Effective_date, pden_area.Expiry_date, pden_area.No_of_gas_wells, pden_area.No_of_injection_wells, pden_area.No_of_oil_wells, pden_area.Ppdm_guid, pden_area.Remark, pden_area.Row_changed_by, pden_area.Row_changed_date, pden_area.Row_created_by, pden_area.Row_created_date, pden_area.Row_effective_date, pden_area.Row_expiry_date, pden_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenArea(c *fiber.Ctx) error {
	var pden_area dto.Pden_area

	setDefaults(&pden_area)

	if err := json.Unmarshal(c.Body(), &pden_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_area.Pden_subtype = id

    if pden_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_area set pden_id = :1, pden_source = :2, area_id = :3, area_type = :4, active_ind = :5, effective_date = :6, expiry_date = :7, no_of_gas_wells = :8, no_of_injection_wells = :9, no_of_oil_wells = :10, ppdm_guid = :11, remark = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where pden_subtype = :20")
	if err != nil {
		return err
	}

	pden_area.Row_changed_date = formatDate(pden_area.Row_changed_date)
	pden_area.Effective_date = formatDateString(pden_area.Effective_date)
	pden_area.Expiry_date = formatDateString(pden_area.Expiry_date)
	pden_area.Row_effective_date = formatDateString(pden_area.Row_effective_date)
	pden_area.Row_expiry_date = formatDateString(pden_area.Row_expiry_date)






	rows, err := stmt.Exec(pden_area.Pden_id, pden_area.Pden_source, pden_area.Area_id, pden_area.Area_type, pden_area.Active_ind, pden_area.Effective_date, pden_area.Expiry_date, pden_area.No_of_gas_wells, pden_area.No_of_injection_wells, pden_area.No_of_oil_wells, pden_area.Ppdm_guid, pden_area.Remark, pden_area.Row_changed_by, pden_area.Row_changed_date, pden_area.Row_created_by, pden_area.Row_effective_date, pden_area.Row_expiry_date, pden_area.Row_quality, pden_area.Pden_subtype)
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

func PatchPdenArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_area set "
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
	query += " where pden_subtype = :id"

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

func DeletePdenArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_area dto.Pden_area
	pden_area.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_area where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_area.Pden_subtype)
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


