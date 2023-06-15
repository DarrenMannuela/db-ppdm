package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaCrew(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_crew")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_crew

	for rows.Next() {
		var ba_crew dto.Ba_crew
		if err := rows.Scan(&ba_crew.Crew_company_ba_id, &ba_crew.Crew_id, &ba_crew.Active_ind, &ba_crew.Cost_per_unit, &ba_crew.Cost_per_unit_currency_uom_uom, &ba_crew.Cost_per_unit_uom, &ba_crew.Crew_abbreviation, &ba_crew.Crew_long_name, &ba_crew.Crew_short_name, &ba_crew.Crew_type, &ba_crew.Effective_date, &ba_crew.Expiry_date, &ba_crew.Ppdm_guid, &ba_crew.Remark, &ba_crew.Source, &ba_crew.Row_changed_by, &ba_crew.Row_changed_date, &ba_crew.Row_created_by, &ba_crew.Row_created_date, &ba_crew.Row_effective_date, &ba_crew.Row_expiry_date, &ba_crew.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_crew to result
		result = append(result, ba_crew)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaCrew(c *fiber.Ctx) error {
	var ba_crew dto.Ba_crew

	setDefaults(&ba_crew)

	if err := json.Unmarshal(c.Body(), &ba_crew); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_crew values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	ba_crew.Row_created_date = formatDate(ba_crew.Row_created_date)
	ba_crew.Row_changed_date = formatDate(ba_crew.Row_changed_date)
	ba_crew.Effective_date = formatDateString(ba_crew.Effective_date)
	ba_crew.Expiry_date = formatDateString(ba_crew.Expiry_date)
	ba_crew.Row_effective_date = formatDateString(ba_crew.Row_effective_date)
	ba_crew.Row_expiry_date = formatDateString(ba_crew.Row_expiry_date)






	rows, err := stmt.Exec(ba_crew.Crew_company_ba_id, ba_crew.Crew_id, ba_crew.Active_ind, ba_crew.Cost_per_unit, ba_crew.Cost_per_unit_currency_uom_uom, ba_crew.Cost_per_unit_uom, ba_crew.Crew_abbreviation, ba_crew.Crew_long_name, ba_crew.Crew_short_name, ba_crew.Crew_type, ba_crew.Effective_date, ba_crew.Expiry_date, ba_crew.Ppdm_guid, ba_crew.Remark, ba_crew.Source, ba_crew.Row_changed_by, ba_crew.Row_changed_date, ba_crew.Row_created_by, ba_crew.Row_created_date, ba_crew.Row_effective_date, ba_crew.Row_expiry_date, ba_crew.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaCrew(c *fiber.Ctx) error {
	var ba_crew dto.Ba_crew

	setDefaults(&ba_crew)

	if err := json.Unmarshal(c.Body(), &ba_crew); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_crew.Crew_company_ba_id = id

    if ba_crew.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_crew set crew_id = :1, active_ind = :2, cost_per_unit = :3, cost_per_unit_currency_uom_uom = :4, cost_per_unit_uom = :5, crew_abbreviation = :6, crew_long_name = :7, crew_short_name = :8, crew_type = :9, effective_date = :10, expiry_date = :11, ppdm_guid = :12, remark = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where crew_company_ba_id = :22")
	if err != nil {
		return err
	}

	ba_crew.Row_changed_date = formatDate(ba_crew.Row_changed_date)
	ba_crew.Effective_date = formatDateString(ba_crew.Effective_date)
	ba_crew.Expiry_date = formatDateString(ba_crew.Expiry_date)
	ba_crew.Row_effective_date = formatDateString(ba_crew.Row_effective_date)
	ba_crew.Row_expiry_date = formatDateString(ba_crew.Row_expiry_date)






	rows, err := stmt.Exec(ba_crew.Crew_id, ba_crew.Active_ind, ba_crew.Cost_per_unit, ba_crew.Cost_per_unit_currency_uom_uom, ba_crew.Cost_per_unit_uom, ba_crew.Crew_abbreviation, ba_crew.Crew_long_name, ba_crew.Crew_short_name, ba_crew.Crew_type, ba_crew.Effective_date, ba_crew.Expiry_date, ba_crew.Ppdm_guid, ba_crew.Remark, ba_crew.Source, ba_crew.Row_changed_by, ba_crew.Row_changed_date, ba_crew.Row_created_by, ba_crew.Row_effective_date, ba_crew.Row_expiry_date, ba_crew.Row_quality, ba_crew.Crew_company_ba_id)
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

func PatchBaCrew(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_crew set "
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
	query += " where crew_company_ba_id = :id"

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

func DeleteBaCrew(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_crew dto.Ba_crew
	ba_crew.Crew_company_ba_id = id

	stmt, err := db.Prepare("delete from ba_crew where crew_company_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_crew.Crew_company_ba_id)
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


