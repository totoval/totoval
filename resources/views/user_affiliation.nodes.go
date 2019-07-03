package views

import (
	"github.com/totoval/framework/view"
)

func init() {
	view.AddView("user_affiliation.nodes", `
{{define "user_affiliation.nodes"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Nodes</title>
</head>
<body>
<div id="main" style="width: 1960px;height:1080px;"></div>
</div>


<script src="https://cdnjs.cloudflare.com/ajax/libs/echarts/4.2.1/echarts.js"></script>
<script>
    var data = {{ .data }}
</script>
<script>
    var myChart = echarts.init(document.getElementById('main'));

    myChart.showLoading();

    myChart.hideLoading();

    myChart.setOption(option = {
        tooltip: {
            trigger: 'item',
            triggerOn: 'mousemove'
        },
        series: [
            {
                type: 'tree',

                data: [data],

                left: '2%',
                right: '2%',
                top: '8%',
                bottom: '20%',

                symbol: 'emptyCircle',

                orient: 'vertical',

                expandAndCollapse: true,

                label: {
                    normal: {
                        position: 'top',
                        rotate: -90,
                        verticalAlign: 'middle',
                        align: 'right',
                        fontSize: 9
                    }
                },

                leaves: {
                    label: {
                        normal: {
                            position: 'bottom',
                            rotate: -90,
                            verticalAlign: 'middle',
                            align: 'left'
                        }
                    }
                },

                animationDurationUpdate: 750
            }
        ]
    });

</script>
</body>
</html>
{{ end }}
`)
}
