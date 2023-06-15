package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFossilTaxonLeaf(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fossil_taxon_leaf")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fossil_taxon_leaf

	for rows.Next() {
		var fossil_taxon_leaf dto.Fossil_taxon_leaf
		if err := rows.Scan(&fossil_taxon_leaf.Taxon_leaf_id, &fossil_taxon_leaf.Active_ind, &fossil_taxon_leaf.Effective_date, &fossil_taxon_leaf.Expiry_date, &fossil_taxon_leaf.Ppdm_guid, &fossil_taxon_leaf.Preferred_ind, &fossil_taxon_leaf.Remark, &fossil_taxon_leaf.Source, &fossil_taxon_leaf.Source_document_id, &fossil_taxon_leaf.Taxon_level, &fossil_taxon_leaf.Taxon_name, &fossil_taxon_leaf.Row_changed_by, &fossil_taxon_leaf.Row_changed_date, &fossil_taxon_leaf.Row_created_by, &fossil_taxon_leaf.Row_created_date, &fossil_taxon_leaf.Row_effective_date, &fossil_taxon_leaf.Row_expiry_date, &fossil_taxon_leaf.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fossil_taxon_leaf to result
		result = append(result, fossil_taxon_leaf)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFossilTaxonLeaf(c *fiber.Ctx) error {
	var fossil_taxon_leaf dto.Fossil_taxon_leaf

	setDefaults(&fossil_taxon_leaf)

	if err := json.Unmarshal(c.Body(), &fossil_taxon_leaf); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fossil_taxon_leaf values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	fossil_taxon_leaf.Row_created_date = formatDate(fossil_taxon_leaf.Row_created_date)
	fossil_taxon_leaf.Row_changed_date = formatDate(fossil_taxon_leaf.Row_changed_date)
	fossil_taxon_leaf.Effective_date = formatDateString(fossil_taxon_leaf.Effective_date)
	fossil_taxon_leaf.Expiry_date = formatDateString(fossil_taxon_leaf.Expiry_date)
	fossil_taxon_leaf.Row_effective_date = formatDateString(fossil_taxon_leaf.Row_effective_date)
	fossil_taxon_leaf.Row_expiry_date = formatDateString(fossil_taxon_leaf.Row_expiry_date)






	rows, err := stmt.Exec(fossil_taxon_leaf.Taxon_leaf_id, fossil_taxon_leaf.Active_ind, fossil_taxon_leaf.Effective_date, fossil_taxon_leaf.Expiry_date, fossil_taxon_leaf.Ppdm_guid, fossil_taxon_leaf.Preferred_ind, fossil_taxon_leaf.Remark, fossil_taxon_leaf.Source, fossil_taxon_leaf.Source_document_id, fossil_taxon_leaf.Taxon_level, fossil_taxon_leaf.Taxon_name, fossil_taxon_leaf.Row_changed_by, fossil_taxon_leaf.Row_changed_date, fossil_taxon_leaf.Row_created_by, fossil_taxon_leaf.Row_created_date, fossil_taxon_leaf.Row_effective_date, fossil_taxon_leaf.Row_expiry_date, fossil_taxon_leaf.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFossilTaxonLeaf(c *fiber.Ctx) error {
	var fossil_taxon_leaf dto.Fossil_taxon_leaf

	setDefaults(&fossil_taxon_leaf)

	if err := json.Unmarshal(c.Body(), &fossil_taxon_leaf); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fossil_taxon_leaf.Taxon_leaf_id = id

    if fossil_taxon_leaf.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fossil_taxon_leaf set active_ind = :1, effective_date = :2, expiry_date = :3, ppdm_guid = :4, preferred_ind = :5, remark = :6, source = :7, source_document_id = :8, taxon_level = :9, taxon_name = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where taxon_leaf_id = :18")
	if err != nil {
		return err
	}

	fossil_taxon_leaf.Row_changed_date = formatDate(fossil_taxon_leaf.Row_changed_date)
	fossil_taxon_leaf.Effective_date = formatDateString(fossil_taxon_leaf.Effective_date)
	fossil_taxon_leaf.Expiry_date = formatDateString(fossil_taxon_leaf.Expiry_date)
	fossil_taxon_leaf.Row_effective_date = formatDateString(fossil_taxon_leaf.Row_effective_date)
	fossil_taxon_leaf.Row_expiry_date = formatDateString(fossil_taxon_leaf.Row_expiry_date)






	rows, err := stmt.Exec(fossil_taxon_leaf.Active_ind, fossil_taxon_leaf.Effective_date, fossil_taxon_leaf.Expiry_date, fossil_taxon_leaf.Ppdm_guid, fossil_taxon_leaf.Preferred_ind, fossil_taxon_leaf.Remark, fossil_taxon_leaf.Source, fossil_taxon_leaf.Source_document_id, fossil_taxon_leaf.Taxon_level, fossil_taxon_leaf.Taxon_name, fossil_taxon_leaf.Row_changed_by, fossil_taxon_leaf.Row_changed_date, fossil_taxon_leaf.Row_created_by, fossil_taxon_leaf.Row_effective_date, fossil_taxon_leaf.Row_expiry_date, fossil_taxon_leaf.Row_quality, fossil_taxon_leaf.Taxon_leaf_id)
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

func PatchFossilTaxonLeaf(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fossil_taxon_leaf set "
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
	query += " where taxon_leaf_id = :id"

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

func DeleteFossilTaxonLeaf(c *fiber.Ctx) error {
	id := c.Params("id")
	var fossil_taxon_leaf dto.Fossil_taxon_leaf
	fossil_taxon_leaf.Taxon_leaf_id = id

	stmt, err := db.Prepare("delete from fossil_taxon_leaf where taxon_leaf_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fossil_taxon_leaf.Taxon_leaf_id)
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


