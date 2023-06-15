package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelMB(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_m_b")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_m_b

	for rows.Next() {
		var sp_parcel_m_b dto.Sp_parcel_m_b
		if err := rows.Scan(&sp_parcel_m_b.Spatial_description_id, &sp_parcel_m_b.Spatial_obs_no, &sp_parcel_m_b.M_b_id, &sp_parcel_m_b.Active_ind, &sp_parcel_m_b.Description, &sp_parcel_m_b.Dls_road_allowance_id, &sp_parcel_m_b.Effective_date, &sp_parcel_m_b.Ew_direction, &sp_parcel_m_b.Ew_distance, &sp_parcel_m_b.Ew_distance_ouom, &sp_parcel_m_b.Ew_start_line, &sp_parcel_m_b.Expiry_date, &sp_parcel_m_b.Inactivation_date, &sp_parcel_m_b.Local_coord_system_id, &sp_parcel_m_b.Location_type, &sp_parcel_m_b.Ns_direction, &sp_parcel_m_b.Ns_distance, &sp_parcel_m_b.Ns_distance_ouom, &sp_parcel_m_b.Ns_start_line, &sp_parcel_m_b.Orientation, &sp_parcel_m_b.Parcel_carter_id, &sp_parcel_m_b.Parcel_congress_id, &sp_parcel_m_b.Parcel_dls_id, &sp_parcel_m_b.Parcel_fps_id, &sp_parcel_m_b.Parcel_libya_id, &sp_parcel_m_b.Parcel_lot_id, &sp_parcel_m_b.Parcel_lot_num, &sp_parcel_m_b.Parcel_lot_type, &sp_parcel_m_b.Parcel_ne_loc_id, &sp_parcel_m_b.Parcel_north_sea_id, &sp_parcel_m_b.Parcel_nts_id, &sp_parcel_m_b.Parcel_offshore_id, &sp_parcel_m_b.Parcel_ohio_id, &sp_parcel_m_b.Parcel_pbl_id, &sp_parcel_m_b.Parcel_texas_id, &sp_parcel_m_b.Ppdm_guid, &sp_parcel_m_b.Reference_loc, &sp_parcel_m_b.Remark, &sp_parcel_m_b.Source, &sp_parcel_m_b.Surface_loc, &sp_parcel_m_b.Row_changed_by, &sp_parcel_m_b.Row_changed_date, &sp_parcel_m_b.Row_created_by, &sp_parcel_m_b.Row_created_date, &sp_parcel_m_b.Row_effective_date, &sp_parcel_m_b.Row_expiry_date, &sp_parcel_m_b.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_m_b to result
		result = append(result, sp_parcel_m_b)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelMB(c *fiber.Ctx) error {
	var sp_parcel_m_b dto.Sp_parcel_m_b

	setDefaults(&sp_parcel_m_b)

	if err := json.Unmarshal(c.Body(), &sp_parcel_m_b); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_m_b values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47)")
	if err != nil {
		return err
	}
	sp_parcel_m_b.Row_created_date = formatDate(sp_parcel_m_b.Row_created_date)
	sp_parcel_m_b.Row_changed_date = formatDate(sp_parcel_m_b.Row_changed_date)
	sp_parcel_m_b.Effective_date = formatDateString(sp_parcel_m_b.Effective_date)
	sp_parcel_m_b.Expiry_date = formatDateString(sp_parcel_m_b.Expiry_date)
	sp_parcel_m_b.Inactivation_date = formatDateString(sp_parcel_m_b.Inactivation_date)
	sp_parcel_m_b.Row_effective_date = formatDateString(sp_parcel_m_b.Row_effective_date)
	sp_parcel_m_b.Row_expiry_date = formatDateString(sp_parcel_m_b.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_m_b.Spatial_description_id, sp_parcel_m_b.Spatial_obs_no, sp_parcel_m_b.M_b_id, sp_parcel_m_b.Active_ind, sp_parcel_m_b.Description, sp_parcel_m_b.Dls_road_allowance_id, sp_parcel_m_b.Effective_date, sp_parcel_m_b.Ew_direction, sp_parcel_m_b.Ew_distance, sp_parcel_m_b.Ew_distance_ouom, sp_parcel_m_b.Ew_start_line, sp_parcel_m_b.Expiry_date, sp_parcel_m_b.Inactivation_date, sp_parcel_m_b.Local_coord_system_id, sp_parcel_m_b.Location_type, sp_parcel_m_b.Ns_direction, sp_parcel_m_b.Ns_distance, sp_parcel_m_b.Ns_distance_ouom, sp_parcel_m_b.Ns_start_line, sp_parcel_m_b.Orientation, sp_parcel_m_b.Parcel_carter_id, sp_parcel_m_b.Parcel_congress_id, sp_parcel_m_b.Parcel_dls_id, sp_parcel_m_b.Parcel_fps_id, sp_parcel_m_b.Parcel_libya_id, sp_parcel_m_b.Parcel_lot_id, sp_parcel_m_b.Parcel_lot_num, sp_parcel_m_b.Parcel_lot_type, sp_parcel_m_b.Parcel_ne_loc_id, sp_parcel_m_b.Parcel_north_sea_id, sp_parcel_m_b.Parcel_nts_id, sp_parcel_m_b.Parcel_offshore_id, sp_parcel_m_b.Parcel_ohio_id, sp_parcel_m_b.Parcel_pbl_id, sp_parcel_m_b.Parcel_texas_id, sp_parcel_m_b.Ppdm_guid, sp_parcel_m_b.Reference_loc, sp_parcel_m_b.Remark, sp_parcel_m_b.Source, sp_parcel_m_b.Surface_loc, sp_parcel_m_b.Row_changed_by, sp_parcel_m_b.Row_changed_date, sp_parcel_m_b.Row_created_by, sp_parcel_m_b.Row_created_date, sp_parcel_m_b.Row_effective_date, sp_parcel_m_b.Row_expiry_date, sp_parcel_m_b.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelMB(c *fiber.Ctx) error {
	var sp_parcel_m_b dto.Sp_parcel_m_b

	setDefaults(&sp_parcel_m_b)

	if err := json.Unmarshal(c.Body(), &sp_parcel_m_b); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_m_b.Spatial_description_id = id

    if sp_parcel_m_b.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_m_b set spatial_obs_no = :1, m_b_id = :2, active_ind = :3, description = :4, dls_road_allowance_id = :5, effective_date = :6, ew_direction = :7, ew_distance = :8, ew_distance_ouom = :9, ew_start_line = :10, expiry_date = :11, inactivation_date = :12, local_coord_system_id = :13, location_type = :14, ns_direction = :15, ns_distance = :16, ns_distance_ouom = :17, ns_start_line = :18, orientation = :19, parcel_carter_id = :20, parcel_congress_id = :21, parcel_dls_id = :22, parcel_fps_id = :23, parcel_libya_id = :24, parcel_lot_id = :25, parcel_lot_num = :26, parcel_lot_type = :27, parcel_ne_loc_id = :28, parcel_north_sea_id = :29, parcel_nts_id = :30, parcel_offshore_id = :31, parcel_ohio_id = :32, parcel_pbl_id = :33, parcel_texas_id = :34, ppdm_guid = :35, reference_loc = :36, remark = :37, source = :38, surface_loc = :39, row_changed_by = :40, row_changed_date = :41, row_created_by = :42, row_effective_date = :43, row_expiry_date = :44, row_quality = :45 where spatial_description_id = :47")
	if err != nil {
		return err
	}

	sp_parcel_m_b.Row_changed_date = formatDate(sp_parcel_m_b.Row_changed_date)
	sp_parcel_m_b.Effective_date = formatDateString(sp_parcel_m_b.Effective_date)
	sp_parcel_m_b.Expiry_date = formatDateString(sp_parcel_m_b.Expiry_date)
	sp_parcel_m_b.Inactivation_date = formatDateString(sp_parcel_m_b.Inactivation_date)
	sp_parcel_m_b.Row_effective_date = formatDateString(sp_parcel_m_b.Row_effective_date)
	sp_parcel_m_b.Row_expiry_date = formatDateString(sp_parcel_m_b.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_m_b.Spatial_obs_no, sp_parcel_m_b.M_b_id, sp_parcel_m_b.Active_ind, sp_parcel_m_b.Description, sp_parcel_m_b.Dls_road_allowance_id, sp_parcel_m_b.Effective_date, sp_parcel_m_b.Ew_direction, sp_parcel_m_b.Ew_distance, sp_parcel_m_b.Ew_distance_ouom, sp_parcel_m_b.Ew_start_line, sp_parcel_m_b.Expiry_date, sp_parcel_m_b.Inactivation_date, sp_parcel_m_b.Local_coord_system_id, sp_parcel_m_b.Location_type, sp_parcel_m_b.Ns_direction, sp_parcel_m_b.Ns_distance, sp_parcel_m_b.Ns_distance_ouom, sp_parcel_m_b.Ns_start_line, sp_parcel_m_b.Orientation, sp_parcel_m_b.Parcel_carter_id, sp_parcel_m_b.Parcel_congress_id, sp_parcel_m_b.Parcel_dls_id, sp_parcel_m_b.Parcel_fps_id, sp_parcel_m_b.Parcel_libya_id, sp_parcel_m_b.Parcel_lot_id, sp_parcel_m_b.Parcel_lot_num, sp_parcel_m_b.Parcel_lot_type, sp_parcel_m_b.Parcel_ne_loc_id, sp_parcel_m_b.Parcel_north_sea_id, sp_parcel_m_b.Parcel_nts_id, sp_parcel_m_b.Parcel_offshore_id, sp_parcel_m_b.Parcel_ohio_id, sp_parcel_m_b.Parcel_pbl_id, sp_parcel_m_b.Parcel_texas_id, sp_parcel_m_b.Ppdm_guid, sp_parcel_m_b.Reference_loc, sp_parcel_m_b.Remark, sp_parcel_m_b.Source, sp_parcel_m_b.Surface_loc, sp_parcel_m_b.Row_changed_by, sp_parcel_m_b.Row_changed_date, sp_parcel_m_b.Row_created_by, sp_parcel_m_b.Row_effective_date, sp_parcel_m_b.Row_expiry_date, sp_parcel_m_b.Row_quality, sp_parcel_m_b.Spatial_description_id)
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

func PatchSpParcelMB(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_m_b set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "inactivation_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where spatial_description_id = :id"

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

func DeleteSpParcelMB(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_m_b dto.Sp_parcel_m_b
	sp_parcel_m_b.Spatial_description_id = id

	stmt, err := db.Prepare("delete from sp_parcel_m_b where spatial_description_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_m_b.Spatial_description_id)
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


