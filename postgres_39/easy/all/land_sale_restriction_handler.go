package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSaleRestriction(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_sale_restriction")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_sale_restriction

	for rows.Next() {
		var land_sale_restriction dto.Land_sale_restriction
		if err := rows.Scan(&land_sale_restriction.Restriction_id, &land_sale_restriction.Restriction_version, &land_sale_restriction.Jurisdiction, &land_sale_restriction.Land_sale_number, &land_sale_restriction.Land_sale_offering_id, &land_sale_restriction.Active_ind, &land_sale_restriction.Effective_date, &land_sale_restriction.Expiry_date, &land_sale_restriction.Ppdm_guid, &land_sale_restriction.Remark, &land_sale_restriction.Source, &land_sale_restriction.Row_changed_by, &land_sale_restriction.Row_changed_date, &land_sale_restriction.Row_created_by, &land_sale_restriction.Row_created_date, &land_sale_restriction.Row_effective_date, &land_sale_restriction.Row_expiry_date, &land_sale_restriction.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_sale_restriction to result
		result = append(result, land_sale_restriction)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSaleRestriction(c *fiber.Ctx) error {
	var land_sale_restriction dto.Land_sale_restriction

	setDefaults(&land_sale_restriction)

	if err := json.Unmarshal(c.Body(), &land_sale_restriction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_sale_restriction values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	land_sale_restriction.Row_created_date = formatDate(land_sale_restriction.Row_created_date)
	land_sale_restriction.Row_changed_date = formatDate(land_sale_restriction.Row_changed_date)
	land_sale_restriction.Effective_date = formatDateString(land_sale_restriction.Effective_date)
	land_sale_restriction.Expiry_date = formatDateString(land_sale_restriction.Expiry_date)
	land_sale_restriction.Row_effective_date = formatDateString(land_sale_restriction.Row_effective_date)
	land_sale_restriction.Row_expiry_date = formatDateString(land_sale_restriction.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_restriction.Restriction_id, land_sale_restriction.Restriction_version, land_sale_restriction.Jurisdiction, land_sale_restriction.Land_sale_number, land_sale_restriction.Land_sale_offering_id, land_sale_restriction.Active_ind, land_sale_restriction.Effective_date, land_sale_restriction.Expiry_date, land_sale_restriction.Ppdm_guid, land_sale_restriction.Remark, land_sale_restriction.Source, land_sale_restriction.Row_changed_by, land_sale_restriction.Row_changed_date, land_sale_restriction.Row_created_by, land_sale_restriction.Row_created_date, land_sale_restriction.Row_effective_date, land_sale_restriction.Row_expiry_date, land_sale_restriction.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSaleRestriction(c *fiber.Ctx) error {
	var land_sale_restriction dto.Land_sale_restriction

	setDefaults(&land_sale_restriction)

	if err := json.Unmarshal(c.Body(), &land_sale_restriction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_sale_restriction.Restriction_id = id

    if land_sale_restriction.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_sale_restriction set restriction_version = :1, jurisdiction = :2, land_sale_number = :3, land_sale_offering_id = :4, active_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where restriction_id = :18")
	if err != nil {
		return err
	}

	land_sale_restriction.Row_changed_date = formatDate(land_sale_restriction.Row_changed_date)
	land_sale_restriction.Effective_date = formatDateString(land_sale_restriction.Effective_date)
	land_sale_restriction.Expiry_date = formatDateString(land_sale_restriction.Expiry_date)
	land_sale_restriction.Row_effective_date = formatDateString(land_sale_restriction.Row_effective_date)
	land_sale_restriction.Row_expiry_date = formatDateString(land_sale_restriction.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_restriction.Restriction_version, land_sale_restriction.Jurisdiction, land_sale_restriction.Land_sale_number, land_sale_restriction.Land_sale_offering_id, land_sale_restriction.Active_ind, land_sale_restriction.Effective_date, land_sale_restriction.Expiry_date, land_sale_restriction.Ppdm_guid, land_sale_restriction.Remark, land_sale_restriction.Source, land_sale_restriction.Row_changed_by, land_sale_restriction.Row_changed_date, land_sale_restriction.Row_created_by, land_sale_restriction.Row_effective_date, land_sale_restriction.Row_expiry_date, land_sale_restriction.Row_quality, land_sale_restriction.Restriction_id)
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

func PatchLandSaleRestriction(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_sale_restriction set "
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
	query += " where restriction_id = :id"

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

func DeleteLandSaleRestriction(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_sale_restriction dto.Land_sale_restriction
	land_sale_restriction.Restriction_id = id

	stmt, err := db.Prepare("delete from land_sale_restriction where restriction_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_sale_restriction.Restriction_id)
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


