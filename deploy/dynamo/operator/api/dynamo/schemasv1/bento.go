/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package schemasv1

import (
	"time"

	"github.com/ai-dynamo/dynamo/deploy/dynamo/operator/api/dynamo/modelschemas"
)

type BentoSchema struct {
	ResourceSchema
	BentoRepositoryUid      string                             `json:"bento_repository_uid"`
	Creator                 *UserSchema                        `json:"creator"`
	Version                 string                             `json:"version"`
	Description             string                             `json:"description"`
	ImageBuildStatus        modelschemas.ImageBuildStatus      `json:"image_build_status"`
	UploadStatus            modelschemas.BentoUploadStatus     `json:"upload_status"`
	UploadStartedAt         *time.Time                         `json:"upload_started_at"`
	UploadFinishedAt        *time.Time                         `json:"upload_finished_at"`
	UploadFinishedReason    string                             `json:"upload_finished_reason"`
	PresignedUploadUrl      string                             `json:"presigned_upload_url"`
	PresignedDownloadUrl    string                             `json:"presigned_download_url"`
	PresignedUrlsDeprecated bool                               `json:"presigned_urls_deprecated"`
	TransmissionStrategy    *modelschemas.TransmissionStrategy `json:"transmission_strategy"`
	UploadId                string                             `json:"upload_id"`
	Manifest                *modelschemas.BentoManifestSchema  `json:"manifest"`
	BuildAt                 time.Time                          `json:"build_at"`
}

type BentoListSchema struct {
	BaseListSchema
	Items []*BentoSchema `json:"items"`
}

type BentoWithRepositorySchema struct {
	BentoSchema
	Repository *BentoRepositorySchema `json:"repository"`
}

type BentoWithRepositoryListSchema struct {
	BaseListSchema
	Items []*BentoWithRepositorySchema `json:"items"`
}

type BentoFullSchema struct {
	BentoWithRepositorySchema
	Models []*ModelWithRepositorySchema `json:"models"`
}

type CreateBentoSchema struct {
	Description string                            `json:"description"`
	Version     string                            `json:"version"`
	Manifest    *modelschemas.BentoManifestSchema `json:"manifest"`
	BuildAt     string                            `json:"build_at"`
	Labels      modelschemas.LabelItemsSchema     `json:"labels"`
}

type UpdateBentoSchema struct {
	Description string                             `json:"description"`
	Version     string                             `json:"version"`
	Manifest    **modelschemas.BentoManifestSchema `json:"manifest"`
	BuildAt     string                             `json:"build_at"`
	Labels      *modelschemas.LabelItemsSchema     `json:"labels,omitempty"`
}

type FinishUploadBentoSchema struct {
	Status *modelschemas.BentoUploadStatus `json:"status"`
	Reason *string                         `json:"reason"`
}
