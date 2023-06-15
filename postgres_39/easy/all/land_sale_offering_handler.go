package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSaleOffering(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_sale_offering")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_sale_offering

	for rows.Next() {
		var land_sale_offering dto.Land_sale_offering
		if err := rows.Scan(&land_sale_offering.Jurisdiction, &land_sale_offering.Land_sale_number, &land_sale_offering.Land_sale_offering_id, &land_sale_offering.Active_ind, &land_sale_offering.Cancel_reason, &land_sale_offering.Effective_date, &land_sale_offering.Expiry_date, &land_sale_offering.Gross_size, &land_sale_offering.Gross_size_ouom, &land_sale_offering.Inactivation_date, &land_sale_offering.Land_offering_status, &land_sale_offering.Land_offering_type, &land_sale_offering.Ppdm_guid, &land_sale_offering.Remark, &land_sale_offering.Source, &land_sale_offering.Status_date, &land_sale_offering.Term_duration, &land_sale_offering.Term_duration_ouom, &land_sale_offering.Row_changed_by, &land_sale_offering.Row_changed_date, &land_sale_offering.Row_created_by, &land_sale_offering.Row_created_date, &land_sale_offering.Row_effective_date, &land_sale_offering.Row_expiry_date, &land_sale_offering.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_sale_offering to result
		result = append(result, land_sale_offering)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSaleOffering(c *fiber.Ctx) error {
	var land_sale_offering dto.Land_sale_offering

	setDefaults(&land_sale_offering)

	if err := json.Unmarshal(c.Body(), &land_sale_offering); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_sale_offering values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	land_sale_offering.Row_created_date = formatDate(land_sale_offering.Row_created_date)
	land_sale_offering.Row_changed_date = formatDate(land_sale_offering.Row_changed_date)
	land_sale_offering.Effective_date = formatDateString(land_sale_offering.Effective_date)
	land_sale_offering.Expiry_date = formatDateString(land_sale_offering.Expiry_date)
	land_sale_offering.Inactivation_date = formatDateString(land_sale_offering.Inactivation_date)
	land_sale_offering.Status_date = formatDateString(land_sale_offering.Status_date)
	land_sale_offering.Row_effective_date = formatDateString(land_sale_offering.Row_effective_date)
	land_sale_offering.Row_expiry_date = formatDateString(land_sale_offering.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_offering.Jurisdiction, land_sale_offering.Land_sale_number, land_sale_offering.Land_sale_offering_id, land_sale_offering.Active_ind, land_sale_offering.Cancel_reason, land_sale_offering.Effective_date, land_sale_offering.Expiry_date, land_sale_offering.Gross_size, land_sale_offering.Gross_size_ouom, land_sale_offering.Inactivation_date, land_sale_offering.Land_offering_status, land_sale_offering.Land_offering_type, land_sale_offering.Ppdm_guid, land_sale_offering.Remark, land_sale_offering.Source, land_sale_offering.Status_date, land_sale_offering.Term_duration, land_sale_offering.Term_duration_ouom, land_sale_offering.Row_changed_by, land_sale_offering.Row_changed_date, land_sale_offering.Row_created_by, land_sale_offering.Row_created_date, land_sale_offering.Row_effective_date, land_sale_offering.Row_expiry_date, land_sale_offering.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSaleOffering(c *fiber.Ctx) error {
	var land_sale_offering dto.Land_sale_offering

	setDefaults(&land_sale_offering)

	if err := json.Unmarshal(c.Body(), &land_sale_offering); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_sale_offering.Jurisdiction = id

    if land_sale_offering.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_sale_offering set land_sale_number = :1, land_sale_offering_id = :2, active_ind = :3, cancel_reason = :4, effective_date = :5, expiry_date = :6, gross_size = :7, gross_size_ouom = :8, inactivation_date = :9, land_offering_status = :10, land_offering_type = :11, ppdm_guid = :12, remark = :13, source = :14, status_date = :15, term_duration = :16, term_duration_ouom = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where jurisdiction = :25")
	if err != nil {
		return err
	}

	land_sale_offering.Row_changed_date = formatDate(land_sale_offering.Row_changed_date)
	land_sale_offering.Effective_date = formatDateString(land_sale_offering.Effective_date)
	land_sale_offering.Expiry_date = formatDateString(land_sale_offering.Expiry_date)
	land_sale_offering.Inactivation_date = formatDateString(land_sale_offering.Inactivation_date)
	land_sale_offering.Status_date = formatDateString(land_sale_offering.Status_date)
	land_sale_offering.Row_effective_date = formatDateString(land_sale_offering.Row_effective_date)
	land_sale_offering.Row_expiry_date = formatDateString(land_sale_offering.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_offering.Land_sale_number, land_sale_offering.Land_sale_offering_id, land_sale_offering.Active_ind, land_sale_offering.Cancel_reason, land_sale_offering.Effective_date, land_sale_offering.Expiry_date, land_sale_offering.Gross_size, land_sale_offering.Gross_size_ouom, land_sale_offering.Inactivation_date, land_sale_offering.Land_offering_status, land_sale_offering.Land_offering_type, land_sale_offering.Ppdm_guid, land_sale_offering.Remark, land_sale_offering.Source, land_sale_offering.Status_date, land_sale_offering.Term_duration, land_sale_offering.Term_duration_ouom, land_sale_offering.Row_changed_by, land_sale_offering.Row_changed_date, land_sale_offering.Row_created_by, land_sale_offering.Row_effective_date, land_sale_offering.Row_expiry_date, land_sale_offering.Row_quality, land_sale_offering.Jurisdiction)
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

func PatchLandSaleOffering(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_sale_offering set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "inactivation_date" || key == "status_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where jurisdiction = :id"

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

func DeleteLandSaleOffering(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_sale_offering dto.Land_sale_offering
	land_sale_offering.Jurisdiction = id

	stmt, err := db.Prepare("delete from land_sale_offering where jurisdiction = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_sale_offering.Jurisdiction)
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


