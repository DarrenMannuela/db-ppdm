package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfBaCrew(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_ba_crew")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_ba_crew

	for rows.Next() {
		var sf_ba_crew dto.Sf_ba_crew
		if err := rows.Scan(&sf_ba_crew.Sf_subtype, &sf_ba_crew.Support_facility_id, &sf_ba_crew.Crew_company_ba_id, &sf_ba_crew.Crew_id, &sf_ba_crew.Sf_ba_crew_obs_no, &sf_ba_crew.Active_ind, &sf_ba_crew.Effective_date, &sf_ba_crew.Expiry_date, &sf_ba_crew.Ppdm_guid, &sf_ba_crew.Remark, &sf_ba_crew.Source, &sf_ba_crew.Row_changed_by, &sf_ba_crew.Row_changed_date, &sf_ba_crew.Row_created_by, &sf_ba_crew.Row_created_date, &sf_ba_crew.Row_effective_date, &sf_ba_crew.Row_expiry_date, &sf_ba_crew.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_ba_crew to result
		result = append(result, sf_ba_crew)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfBaCrew(c *fiber.Ctx) error {
	var sf_ba_crew dto.Sf_ba_crew

	setDefaults(&sf_ba_crew)

	if err := json.Unmarshal(c.Body(), &sf_ba_crew); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_ba_crew values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	sf_ba_crew.Row_created_date = formatDate(sf_ba_crew.Row_created_date)
	sf_ba_crew.Row_changed_date = formatDate(sf_ba_crew.Row_changed_date)
	sf_ba_crew.Effective_date = formatDateString(sf_ba_crew.Effective_date)
	sf_ba_crew.Expiry_date = formatDateString(sf_ba_crew.Expiry_date)
	sf_ba_crew.Row_effective_date = formatDateString(sf_ba_crew.Row_effective_date)
	sf_ba_crew.Row_expiry_date = formatDateString(sf_ba_crew.Row_expiry_date)






	rows, err := stmt.Exec(sf_ba_crew.Sf_subtype, sf_ba_crew.Support_facility_id, sf_ba_crew.Crew_company_ba_id, sf_ba_crew.Crew_id, sf_ba_crew.Sf_ba_crew_obs_no, sf_ba_crew.Active_ind, sf_ba_crew.Effective_date, sf_ba_crew.Expiry_date, sf_ba_crew.Ppdm_guid, sf_ba_crew.Remark, sf_ba_crew.Source, sf_ba_crew.Row_changed_by, sf_ba_crew.Row_changed_date, sf_ba_crew.Row_created_by, sf_ba_crew.Row_created_date, sf_ba_crew.Row_effective_date, sf_ba_crew.Row_expiry_date, sf_ba_crew.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfBaCrew(c *fiber.Ctx) error {
	var sf_ba_crew dto.Sf_ba_crew

	setDefaults(&sf_ba_crew)

	if err := json.Unmarshal(c.Body(), &sf_ba_crew); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_ba_crew.Sf_subtype = id

    if sf_ba_crew.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_ba_crew set support_facility_id = :1, crew_company_ba_id = :2, crew_id = :3, sf_ba_crew_obs_no = :4, active_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where sf_subtype = :18")
	if err != nil {
		return err
	}

	sf_ba_crew.Row_changed_date = formatDate(sf_ba_crew.Row_changed_date)
	sf_ba_crew.Effective_date = formatDateString(sf_ba_crew.Effective_date)
	sf_ba_crew.Expiry_date = formatDateString(sf_ba_crew.Expiry_date)
	sf_ba_crew.Row_effective_date = formatDateString(sf_ba_crew.Row_effective_date)
	sf_ba_crew.Row_expiry_date = formatDateString(sf_ba_crew.Row_expiry_date)






	rows, err := stmt.Exec(sf_ba_crew.Support_facility_id, sf_ba_crew.Crew_company_ba_id, sf_ba_crew.Crew_id, sf_ba_crew.Sf_ba_crew_obs_no, sf_ba_crew.Active_ind, sf_ba_crew.Effective_date, sf_ba_crew.Expiry_date, sf_ba_crew.Ppdm_guid, sf_ba_crew.Remark, sf_ba_crew.Source, sf_ba_crew.Row_changed_by, sf_ba_crew.Row_changed_date, sf_ba_crew.Row_created_by, sf_ba_crew.Row_effective_date, sf_ba_crew.Row_expiry_date, sf_ba_crew.Row_quality, sf_ba_crew.Sf_subtype)
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

func PatchSfBaCrew(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_ba_crew set "
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
	query += " where sf_subtype = :id"

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

func DeleteSfBaCrew(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_ba_crew dto.Sf_ba_crew
	sf_ba_crew.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_ba_crew where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_ba_crew.Sf_subtype)
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


