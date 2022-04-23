// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package rbac

import "github.com/storyicon/grbac"

func Rules() grbac.Rules {
	return grbac.Rules{
		{
			ID: 0, // 越大，优先级最高
			Resource: &grbac.Resource{
				Host:   "*",
				Path:   "**",
				Method: "*",
			},
			Permission: &grbac.Permission{
				AuthorizedRoles: []string{"*"},
				AllowAnyone:     true,
			},
		},
		{
			ID: 100, // 越大，优先级最高
			Resource: &grbac.Resource{
				Host:   "*",
				Path:   "/swagger/**",
				Method: "*",
			},
			Permission: &grbac.Permission{
				AuthorizedRoles: []string{"*"},
				AllowAnyone:     true,
			},
		},
		{
			ID: 100, // 越大，优先级最高
			Resource: &grbac.Resource{
				Host:   "/api/**",
				Path:   "**",
				Method: "*",
			},
			Permission: &grbac.Permission{
				AuthorizedRoles: []string{"*"},
			},
		},
		{
			ID: 100, // 越大，优先级最高
			Resource: &grbac.Resource{
				Host:   "/apis/**",
				Path:   "**",
				Method: "*",
			},
			Permission: &grbac.Permission{
				AuthorizedRoles: []string{"*"},
			},
		},
	}
}
