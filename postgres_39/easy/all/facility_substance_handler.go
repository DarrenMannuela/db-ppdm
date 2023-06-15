package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilitySubstance(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_substance")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_substance

	for rows.Next() {
		var facility_substance dto.Facility_substance
		if err := rows.Scan(&facility_substance.Facility_id, &facility_substance.Facility_type, &facility_substance.Substance_id, &facility_substance.Active_ind, &facility_substance.Average_capacity, &facility_substance.Capacity_ouom, &facility_substance.Capacity_uom, &facility_substance.Effective_date, &facility_substance.Expiry_date, &facility_substance.Max_capacity, &facility_substance.Ppdm_guid, &facility_substance.Remark, &facility_substance.Source, &facility_substance.Strat_name_set_id, &facility_substance.Strat_unit_id, &facility_substance.Substance_excluded_ind, &facility_substance.Substance_included_ind, &facility_substance.Row_changed_by, &facility_substance.Row_changed_date, &facility_substance.Row_created_by, &facility_substance.Row_created_date, &facility_substance.Row_effective_date, &facility_substance.Row_expiry_date, &facility_substance.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_substance to result
		result = append(result, facility_substance)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilitySubstance(c *fiber.Ctx) error {
	var facility_substance dto.Facility_substance

	setDefaults(&facility_substance)

	if err := json.Unmarshal(c.Body(), &facility_substance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_substance values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	facility_substance.Row_created_date = formatDate(facility_substance.Row_created_date)
	facility_substance.Row_changed_date = formatDate(facility_substance.Row_changed_date)
	facility_substance.Effective_date = formatDateString(facility_substance.Effective_date)
	facility_substance.Expiry_date = formatDateString(facility_substance.Expiry_date)
	facility_substance.Row_effective_date = formatDateString(facility_substance.Row_effective_date)
	facility_substance.Row_expiry_date = formatDateString(facility_substance.Row_expiry_date)






	rows, err := stmt.Exec(facility_substance.Facility_id, facility_substance.Facility_type, facility_substance.Substance_id, facility_substance.Active_ind, facility_substance.Average_capacity, facility_substance.Capacity_ouom, facility_substance.Capacity_uom, facility_substance.Effective_date, facility_substance.Expiry_date, facility_substance.Max_capacity, facility_substance.Ppdm_guid, facility_substance.Remark, facility_substance.Source, facility_substance.Strat_name_set_id, facility_substance.Strat_unit_id, facility_substance.Substance_excluded_ind, facility_substance.Substance_included_ind, facility_substance.Row_changed_by, facility_substance.Row_changed_date, facility_substance.Row_created_by, facility_substance.Row_created_date, facility_substance.Row_effective_date, facility_substance.Row_expiry_date, facility_substance.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilitySubstance(c *fiber.Ctx) error {
	var facility_substance dto.Facility_substance

	setDefaults(&facility_substance)

	if err := json.Unmarshal(c.Body(), &facility_substance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_substance.Facility_id = id

    if facility_substance.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_substance set facility_type = :1, substance_id = :2, active_ind = :3, average_capacity = :4, capacity_ouom = :5, capacity_uom = :6, effective_date = :7, expiry_date = :8, max_capacity = :9, ppdm_guid = :10, remark = :11, source = :12, strat_name_set_id = :13, strat_unit_id = :14, substance_excluded_ind = :15, substance_included_ind = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where facility_id = :24")
	if err != nil {
		return err
	}

	facility_substance.Row_changed_date = formatDate(facility_substance.Row_changed_date)
	facility_substance.Effective_date = formatDateString(facility_substance.Effective_date)
	facility_substance.Expiry_date = formatDateString(facility_substance.Expiry_date)
	facility_substance.Row_effective_date = formatDateString(facility_substance.Row_effective_date)
	facility_substance.Row_expiry_date = formatDateString(facility_substance.Row_expiry_date)






	rows, err := stmt.Exec(facility_substance.Facility_type, facility_substance.Substance_id, facility_substance.Active_ind, facility_substance.Average_capacity, facility_substance.Capacity_ouom, facility_substance.Capacity_uom, facility_substance.Effective_date, facility_substance.Expiry_date, facility_substance.Max_capacity, facility_substance.Ppdm_guid, facility_substance.Remark, facility_substance.Source, facility_substance.Strat_name_set_id, facility_substance.Strat_unit_id, facility_substance.Substance_excluded_ind, facility_substance.Substance_included_ind, facility_substance.Row_changed_by, facility_substance.Row_changed_date, facility_substance.Row_created_by, facility_substance.Row_effective_date, facility_substance.Row_expiry_date, facility_substance.Row_quality, facility_substance.Facility_id)
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

func PatchFacilitySubstance(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_substance set "
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

func DeleteFacilitySubstance(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_substance dto.Facility_substance
	facility_substance.Facility_id = id

	stmt, err := db.Prepare("delete from facility_substance where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_substance.Facility_id)
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


