package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFossilAssemblage(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fossil_assemblage")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fossil_assemblage

	for rows.Next() {
		var fossil_assemblage dto.Fossil_assemblage
		if err := rows.Scan(&fossil_assemblage.Strat_name_set_id, &fossil_assemblage.Strat_unit_id, &fossil_assemblage.Taxon_leaf_id, &fossil_assemblage.Interp_id, &fossil_assemblage.Active_ind, &fossil_assemblage.Assemblage_type, &fossil_assemblage.Effective_date, &fossil_assemblage.Expiry_date, &fossil_assemblage.Oldest_ind, &fossil_assemblage.Ppdm_guid, &fossil_assemblage.Primary_marker_ind, &fossil_assemblage.Remark, &fossil_assemblage.Source, &fossil_assemblage.Source_document_id, &fossil_assemblage.Row_changed_by, &fossil_assemblage.Row_changed_date, &fossil_assemblage.Row_created_by, &fossil_assemblage.Row_created_date, &fossil_assemblage.Row_effective_date, &fossil_assemblage.Row_expiry_date, &fossil_assemblage.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fossil_assemblage to result
		result = append(result, fossil_assemblage)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFossilAssemblage(c *fiber.Ctx) error {
	var fossil_assemblage dto.Fossil_assemblage

	setDefaults(&fossil_assemblage)

	if err := json.Unmarshal(c.Body(), &fossil_assemblage); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fossil_assemblage values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	fossil_assemblage.Row_created_date = formatDate(fossil_assemblage.Row_created_date)
	fossil_assemblage.Row_changed_date = formatDate(fossil_assemblage.Row_changed_date)
	fossil_assemblage.Effective_date = formatDateString(fossil_assemblage.Effective_date)
	fossil_assemblage.Expiry_date = formatDateString(fossil_assemblage.Expiry_date)
	fossil_assemblage.Row_effective_date = formatDateString(fossil_assemblage.Row_effective_date)
	fossil_assemblage.Row_expiry_date = formatDateString(fossil_assemblage.Row_expiry_date)






	rows, err := stmt.Exec(fossil_assemblage.Strat_name_set_id, fossil_assemblage.Strat_unit_id, fossil_assemblage.Taxon_leaf_id, fossil_assemblage.Interp_id, fossil_assemblage.Active_ind, fossil_assemblage.Assemblage_type, fossil_assemblage.Effective_date, fossil_assemblage.Expiry_date, fossil_assemblage.Oldest_ind, fossil_assemblage.Ppdm_guid, fossil_assemblage.Primary_marker_ind, fossil_assemblage.Remark, fossil_assemblage.Source, fossil_assemblage.Source_document_id, fossil_assemblage.Row_changed_by, fossil_assemblage.Row_changed_date, fossil_assemblage.Row_created_by, fossil_assemblage.Row_created_date, fossil_assemblage.Row_effective_date, fossil_assemblage.Row_expiry_date, fossil_assemblage.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFossilAssemblage(c *fiber.Ctx) error {
	var fossil_assemblage dto.Fossil_assemblage

	setDefaults(&fossil_assemblage)

	if err := json.Unmarshal(c.Body(), &fossil_assemblage); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fossil_assemblage.Strat_name_set_id = id

    if fossil_assemblage.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fossil_assemblage set strat_unit_id = :1, taxon_leaf_id = :2, interp_id = :3, active_ind = :4, assemblage_type = :5, effective_date = :6, expiry_date = :7, oldest_ind = :8, ppdm_guid = :9, primary_marker_ind = :10, remark = :11, source = :12, source_document_id = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where strat_name_set_id = :21")
	if err != nil {
		return err
	}

	fossil_assemblage.Row_changed_date = formatDate(fossil_assemblage.Row_changed_date)
	fossil_assemblage.Effective_date = formatDateString(fossil_assemblage.Effective_date)
	fossil_assemblage.Expiry_date = formatDateString(fossil_assemblage.Expiry_date)
	fossil_assemblage.Row_effective_date = formatDateString(fossil_assemblage.Row_effective_date)
	fossil_assemblage.Row_expiry_date = formatDateString(fossil_assemblage.Row_expiry_date)






	rows, err := stmt.Exec(fossil_assemblage.Strat_unit_id, fossil_assemblage.Taxon_leaf_id, fossil_assemblage.Interp_id, fossil_assemblage.Active_ind, fossil_assemblage.Assemblage_type, fossil_assemblage.Effective_date, fossil_assemblage.Expiry_date, fossil_assemblage.Oldest_ind, fossil_assemblage.Ppdm_guid, fossil_assemblage.Primary_marker_ind, fossil_assemblage.Remark, fossil_assemblage.Source, fossil_assemblage.Source_document_id, fossil_assemblage.Row_changed_by, fossil_assemblage.Row_changed_date, fossil_assemblage.Row_created_by, fossil_assemblage.Row_effective_date, fossil_assemblage.Row_expiry_date, fossil_assemblage.Row_quality, fossil_assemblage.Strat_name_set_id)
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

func PatchFossilAssemblage(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fossil_assemblage set "
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
	query += " where strat_name_set_id = :id"

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

func DeleteFossilAssemblage(c *fiber.Ctx) error {
	id := c.Params("id")
	var fossil_assemblage dto.Fossil_assemblage
	fossil_assemblage.Strat_name_set_id = id

	stmt, err := db.Prepare("delete from fossil_assemblage where strat_name_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fossil_assemblage.Strat_name_set_id)
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


