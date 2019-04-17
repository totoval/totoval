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
    // var option = {
    //     title: {
    //         text: 'WORLD COFFEE RESEARCH SENSORY LEXICON',
    //         subtext: 'Source: https://worldcoffeeresearch.org/work/sensory-lexicon/',
    //         textStyle: {
    //             fontSize: 14,
    //             align: 'center'
    //         },
    //         subtextStyle: {
    //             align: 'center'
    //         },
    //         sublink: 'https://worldcoffeeresearch.org/work/sensory-lexicon/'
    //     },
    //     series: {
    //         type: 'sunburst',
    //         highlightPolicy: 'ancestor',
    //         data: data.children,
    //         radius: [0, '95%'],
    //         sort: null,
    //         levels: [{}, {
    //             r0: '15%',
    //             r: '35%',
    //             itemStyle: {
    //                 borderWidth: 2
    //             },
    //             label: {
    //                 rotate: 'tangential'
    //             }
    //         }, {
    //             r0: '35%',
    //             r: '70%',
    //             label: {
    //                 align: 'right'
    //             }
    //         }, {
    //             r0: '70%',
    //             r: '72%',
    //             label: {
    //                 position: 'outside',
    //                 padding: 3,
    //                 silent: false
    //             },
    //             itemStyle: {
    //                 borderWidth: 3
    //             }
    //         }]
    //     }
    // };

    // var option = {
    //     series: {
    //         radius: ['15%', '80%'],
    //         type: 'sunburst',
    //         sort: null,
    //         highlightPolicy: 'ancestor',
    //         data: data.children,
    //         label: {
    //             rotate: 'radial'
    //         },
    //         levels: [],
    //         itemStyle: {
    //             color: '#ddd',
    //             borderWidth: 2
    //         }
    //     }
    // };

    //myChart.setOption(option);

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
