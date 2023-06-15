package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSaleRestRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_sale_rest_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_sale_rest_remark

	for rows.Next() {
		var land_sale_rest_remark dto.Land_sale_rest_remark
		if err := rows.Scan(&land_sale_rest_remark.Restriction_id, &land_sale_rest_remark.Restriction_version, &land_sale_rest_remark.Jurisdiction, &land_sale_rest_remark.Land_sale_number, &land_sale_rest_remark.Land_sale_offering_id, &land_sale_rest_remark.Restriction_remark_id, &land_sale_rest_remark.Active_ind, &land_sale_rest_remark.Effective_date, &land_sale_rest_remark.Expiry_date, &land_sale_rest_remark.Ppdm_guid, &land_sale_rest_remark.Remark, &land_sale_rest_remark.Restriction_remark_type, &land_sale_rest_remark.Source, &land_sale_rest_remark.Row_changed_by, &land_sale_rest_remark.Row_changed_date, &land_sale_rest_remark.Row_created_by, &land_sale_rest_remark.Row_created_date, &land_sale_rest_remark.Row_effective_date, &land_sale_rest_remark.Row_expiry_date, &land_sale_rest_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_sale_rest_remark to result
		result = append(result, land_sale_rest_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSaleRestRemark(c *fiber.Ctx) error {
	var land_sale_rest_remark dto.Land_sale_rest_remark

	setDefaults(&land_sale_rest_remark)

	if err := json.Unmarshal(c.Body(), &land_sale_rest_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_sale_rest_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	land_sale_rest_remark.Row_created_date = formatDate(land_sale_rest_remark.Row_created_date)
	land_sale_rest_remark.Row_changed_date = formatDate(land_sale_rest_remark.Row_changed_date)
	land_sale_rest_remark.Effective_date = formatDateString(land_sale_rest_remark.Effective_date)
	land_sale_rest_remark.Expiry_date = formatDateString(land_sale_rest_remark.Expiry_date)
	land_sale_rest_remark.Row_effective_date = formatDateString(land_sale_rest_remark.Row_effective_date)
	land_sale_rest_remark.Row_expiry_date = formatDateString(land_sale_rest_remark.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_rest_remark.Restriction_id, land_sale_rest_remark.Restriction_version, land_sale_rest_remark.Jurisdiction, land_sale_rest_remark.Land_sale_number, land_sale_rest_remark.Land_sale_offering_id, land_sale_rest_remark.Restriction_remark_id, land_sale_rest_remark.Active_ind, land_sale_rest_remark.Effective_date, land_sale_rest_remark.Expiry_date, land_sale_rest_remark.Ppdm_guid, land_sale_rest_remark.Remark, land_sale_rest_remark.Restriction_remark_type, land_sale_rest_remark.Source, land_sale_rest_remark.Row_changed_by, land_sale_rest_remark.Row_changed_date, land_sale_rest_remark.Row_created_by, land_sale_rest_remark.Row_created_date, land_sale_rest_remark.Row_effective_date, land_sale_rest_remark.Row_expiry_date, land_sale_rest_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSaleRestRemark(c *fiber.Ctx) error {
	var land_sale_rest_remark dto.Land_sale_rest_remark

	setDefaults(&land_sale_rest_remark)

	if err := json.Unmarshal(c.Body(), &land_sale_rest_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_sale_rest_remark.Restriction_id = id

    if land_sale_rest_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_sale_rest_remark set restriction_version = :1, jurisdiction = :2, land_sale_number = :3, land_sale_offering_id = :4, restriction_remark_id = :5, active_ind = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, restriction_remark_type = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where restriction_id = :20")
	if err != nil {
		return err
	}

	land_sale_rest_remark.Row_changed_date = formatDate(land_sale_rest_remark.Row_changed_date)
	land_sale_rest_remark.Effective_date = formatDateString(land_sale_rest_remark.Effective_date)
	land_sale_rest_remark.Expiry_date = formatDateString(land_sale_rest_remark.Expiry_date)
	land_sale_rest_remark.Row_effective_date = formatDateString(land_sale_rest_remark.Row_effective_date)
	land_sale_rest_remark.Row_expiry_date = formatDateString(land_sale_rest_remark.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_rest_remark.Restriction_version, land_sale_rest_remark.Jurisdiction, land_sale_rest_remark.Land_sale_number, land_sale_rest_remark.Land_sale_offering_id, land_sale_rest_remark.Restriction_remark_id, land_sale_rest_remark.Active_ind, land_sale_rest_remark.Effective_date, land_sale_rest_remark.Expiry_date, land_sale_rest_remark.Ppdm_guid, land_sale_rest_remark.Remark, land_sale_rest_remark.Restriction_remark_type, land_sale_rest_remark.Source, land_sale_rest_remark.Row_changed_by, land_sale_rest_remark.Row_changed_date, land_sale_rest_remark.Row_created_by, land_sale_rest_remark.Row_effective_date, land_sale_rest_remark.Row_expiry_date, land_sale_rest_remark.Row_quality, land_sale_rest_remark.Restriction_id)
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

func PatchLandSaleRestRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_sale_rest_remark set "
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

func DeleteLandSaleRestRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_sale_rest_remark dto.Land_sale_rest_remark
	land_sale_rest_remark.Restriction_id = id

	stmt, err := db.Prepare("delete from land_sale_rest_remark where restriction_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_sale_rest_remark.Restriction_id)
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


